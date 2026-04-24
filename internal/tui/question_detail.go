package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cicbyte/answer-cli/internal/models"
)

func (m appModel) updateQuestionDetail(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseWheelUp:
			if m.detailScroll > 0 {
				m.detailScroll--
			}
			return m, nil
		case tea.MouseWheelDown:
			m.detailScroll++
			return m, nil
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.activeView = viewQuestionList
			m.question = nil
			m.answers = nil
			m.comments = make(map[string][]models.CommentInfo)
			m.detailScroll = 0
			return m, loadQuestionsCmd(m.cli, m.qPage, m.qPageSize, m.qOrder)
		case "up", "k":
			if m.detailScroll > 0 {
				m.detailScroll--
			}
		case "down", "j":
			m.detailScroll++
		case "v":
			if m.question != nil {
				return m, voteUpCmd(m.cli, m.question.ID)
			}
		case "r":
			if m.question != nil {
				ta := newEditorTextarea("编写你的回答...")
				m.editorActive = true
				m.editorMode = modeAnswer
				m.editorInput = ta
				m.editorTarget = m.question.ID
				return m, ta.Focus()
			}
		case "c":
			if m.question != nil {
				ta := newEditorTextarea("添加评论...")
				m.editorActive = true
				m.editorMode = modeComment
				m.editorInput = ta
				m.editorTarget = m.question.ID
				return m, ta.Focus()
			}
		}
	}
	return m, nil
}

func (m appModel) viewQuestionDetail() string {
	if m.dLoading {
		return styleMuted.Render("加载中...")
	}
	if m.question == nil {
		return styleMuted.Render("无内容")
	}

	header := m.renderDetailHeader()
	body := m.renderDetailBody()
	bodyLines := strings.Split(body, "\n")

	// 可用高度 = 终端高度 - 头部行数 - 状态栏(1)
	headerLines := strings.Count(header, "\n") + 1
	viewH := m.height - headerLines - 1
	if viewH < 1 {
		viewH = 1
	}

	// 限制滚动范围
	maxScroll := max(len(bodyLines)-viewH, 0)
	if m.detailScroll > maxScroll {
		m.detailScroll = maxScroll
	}
	if m.detailScroll < 0 {
		m.detailScroll = 0
	}

	if m.detailScroll >= len(bodyLines) {
		m.detailScroll = len(bodyLines) - 1
	}

	visible := bodyLines[m.detailScroll:]
	if len(visible) > viewH {
		visible = visible[:viewH]
	}

	return header + strings.Join(visible, "\n")
}

func (m appModel) renderDetailHeader() string {
	var b strings.Builder
	w := maxWidth(m.width)

	b.WriteString(styleTitle.Render(m.question.Title))
	b.WriteString("\n")

	author := "匿名"
	if m.question.UserInfo != nil {
		author = m.question.UserInfo.DisplayName
	}
	date := models.FormatTimestamp(m.question.CreatedAt).Format("2006-01-02 15:04")
	votes := styleVote.Render(fmt.Sprintf("↑%d", m.question.VoteCount))

	meta := fmt.Sprintf("%s · %s · %s · 浏览%d · %s", author, date, votes, m.question.ViewCount, statusLabel(m.question.Status))
	if m.question.AcceptedAnswerID != "" {
		meta += " · " + styleAccepted.Render("✓ 已采纳")
	}
	b.WriteString(styleMuted.Render(meta))

	if len(m.question.Tags) > 0 {
		b.WriteString("\n")
		for _, t := range m.question.Tags {
			b.WriteString(" " + styleTag.Render(t.DisplayName))
		}
	}
	b.WriteString("\n")
	b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))

	return b.String()
}

func (m appModel) renderDetailBody() string {
	if m.question == nil {
		return ""
	}

	var b strings.Builder
	w := maxWidth(m.width)

	// 问题内容
	b.WriteString("\n")
	b.WriteString(renderMarkdown(m.question.Content, w))

	// 回答区
	if len(m.answers) == 0 {
		b.WriteString("\n\n")
		b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
		b.WriteString("\n")
		b.WriteString(styleMuted.Render("暂无回答"))
	} else {
		for i, a := range m.answers {
			isAccepted := a.ID == m.question.AcceptedAnswerID
			aAuthor := "匿名"
			if a.UserInfo != nil {
				aAuthor = a.UserInfo.DisplayName
			}
			aDate := models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04")

			b.WriteString("\n\n")
			b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
			b.WriteString("\n")

			title := fmt.Sprintf(" 回答 %d ", i+1)
			if isAccepted {
				title += styleAccepted.Render("✓ 已采纳")
			}
			info := fmt.Sprintf("%s · %s · ↑%d", aAuthor, aDate, a.VoteCount)
			if a.CommentCount > 0 {
				info += fmt.Sprintf(" · 💬%d", a.CommentCount)
			}

			b.WriteString(styleTitle.Render(title))
			if isAccepted {
				b.WriteString(styleAccepted.Render(strings.Repeat(" ", max(w-lipgloss.Width(title+info), 1))))
			}
			b.WriteString("\n")
			b.WriteString(styleMuted.Render(" " + info))
			b.WriteString("\n\n")

			b.WriteString(renderMarkdown(a.Content, w))

			comments := m.comments[a.ID]
			if len(comments) > 0 {
				b.WriteString("\n")
				for _, c := range comments {
					cAuthor := "匿名"
					if c.UserInfo != nil {
						cAuthor = c.UserInfo.DisplayName
					}
					cDate := models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04")
					b.WriteString(styleDim.Render(fmt.Sprintf("  └ %s %s: %s", cAuthor, cDate, c.OriginalText)))
					b.WriteString("\n")
				}
			}
		}
	}

	// 问题评论
	qComments := m.comments[m.question.ID]
	if len(qComments) > 0 {
		b.WriteString("\n\n")
		b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
		b.WriteString("\n")
		b.WriteString(styleMuted.Render("问题评论"))
		b.WriteString("\n")
		for _, c := range qComments {
			cAuthor := "匿名"
			if c.UserInfo != nil {
				cAuthor = c.UserInfo.DisplayName
			}
			cDate := models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04")
			b.WriteString(styleDim.Render(fmt.Sprintf("  └ %s %s: %s", cAuthor, cDate, c.OriginalText)))
			b.WriteString("\n")
		}
	}

	return b.String()
}

func statusLabel(s int) string {
	switch s {
	case 1:
		return "正常"
	case 2:
		return "已关闭"
	case 10:
		return "待审核"
	default:
		return fmt.Sprintf("状态%d", s)
	}
}
