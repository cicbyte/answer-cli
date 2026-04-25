package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cicbyte/answer-cli/internal/models"
)

func (m appModel) updateQuestionDetail(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.detailHeaderLines > 0 {
		if vpH := m.height - m.detailHeaderLines - 2; vpH > 0 {
			m.viewport.Height = vpH
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "esc":
			m.activeView = viewQuestionList
			m.question = nil
			m.answers = nil
			m.comments = make(map[string][]models.CommentInfo)
			m.detailHeaderLines = 0
			return m, loadQuestionsCmd(m.cli, m.qPage, m.qPageSize, m.qOrder)
		case "up", "k":
			m.viewport.LineUp(1)
			return m, nil
		case "down", "j":
			m.viewport.LineDown(1)
			return m, nil
		case "home":
			m.viewport.GotoTop()
			return m, nil
		case "end":
			m.viewport.GotoBottom()
			return m, nil
		}
	case tea.MouseMsg:
		switch {
		case msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonWheelUp:
			m.viewport.LineUp(1)
		case msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonWheelDown:
			m.viewport.LineDown(1)
		}
		return m, nil
	}

	var cmds []tea.Cmd

	if qm, ok := msg.(questionDetailLoadedMsg); ok {
		m.question = qm.question
		m.dLoading = false
		m.detailHeaderLines = strings.Count(m.renderDetailHeader(), "\n") + 1
		if vpH := m.height - m.detailHeaderLines - 2; vpH > 0 {
			m.viewport.Height = vpH
		}
		m.detailContent = m.buildDetailContent()
		m.viewport.SetContent(m.detailContent)
		m.viewport.GotoTop()
		cmds = append(cmds, loadAnswersCmd(m.cli, qm.question.ID))
		if qm.question.ID != "" {
			cmds = append(cmds, loadCommentsCmd(m.cli, qm.question.ID))
		}
	} else if am, ok := msg.(answersLoadedMsg); ok {
		m.answers = am.answers
		m.detailContent = m.buildDetailContent()
		m.viewport.SetContent(m.detailContent)
		for _, a := range am.answers {
			if a.CommentCount > 0 {
				cmds = append(cmds, loadCommentsCmd(m.cli, a.ID))
			}
		}
	} else if cm, ok := msg.(commentsLoadedMsg); ok {
		m.comments[cm.objectID] = cm.comments
		m.detailContent = m.buildDetailContent()
		m.viewport.SetContent(m.detailContent)
	}

	if len(cmds) > 0 {
		return m, tea.Batch(cmds...)
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
	body := m.viewport.View()

	return header + "\n" + body
}

func (m appModel) buildDetailContent() string {
	if m.question == nil {
		return ""
	}

	var b strings.Builder
	w := maxWidth(m.width)

	b.WriteString(renderMarkdown(m.question.Content, w))

	qComments := m.comments[m.question.ID]
	if len(qComments) > 0 {
		b.WriteString("\n\n")
		b.WriteString(styleMuted.Render("问题评论"))
		b.WriteString("\n")
		for _, c := range qComments {
			cAuthor := "匿名"
			if c.DisplayAuthor() != "匿名" {
				cAuthor = c.DisplayAuthor()
				if cAuthor == "" {
					cAuthor = c.Username
				}
			}
			cDate := models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04")
			b.WriteString(styleDim.Render(fmt.Sprintf("  └ #%s %s %s: %s", c.CommentID, cAuthor, cDate, c.OriginalText)))
			b.WriteString("\n")
		}
	}

	if len(m.answers) == 0 {
		b.WriteString("\n\n")
		b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
		b.WriteString("\n")
		b.WriteString(styleMuted.Render("暂无回答"))
	} else {
		for _, a := range m.answers {
			isAccepted := a.ID == m.question.AcceptedAnswerID
			aAuthor := "匿名"
			if a.UserInfo != nil {
				aAuthor = a.UserInfo.DisplayName
			}
			aDate := models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04")

			b.WriteString("\n\n")
			b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
			b.WriteString("\n")

			title := fmt.Sprintf(" 回答 #%s ", a.ID)
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
					if c.DisplayAuthor() != "匿名" {
						cAuthor = c.DisplayAuthor()
						if cAuthor == "" {
							cAuthor = c.Username
						}
					}
					cDate := models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04")
					b.WriteString(styleDim.Render(fmt.Sprintf("  └ #%s %s %s: %s", c.CommentID, cAuthor, cDate, c.OriginalText)))
					b.WriteString("\n")
				}
			}
		}
	}

	return b.String()
}

func (m appModel) renderDetailHeader() string {
	var b strings.Builder
	w := maxWidth(m.width)

	b.WriteString(styleTitle.Render(m.question.Title))
	b.WriteString("\n")
	b.WriteString(styleDim.Render(fmt.Sprintf("#%s", m.question.ID)))
	b.WriteString("\n")

	author := "匿名"
	if m.question.UserInfo != nil {
		author = m.question.UserInfo.DisplayName
	}
	date := models.FormatTimestamp(m.question.CreatedAt).Format("2006-01-02 15:04")
	votes := styleVote.Render(fmt.Sprintf("↑%d", m.question.VoteCount))

	meta := fmt.Sprintf("%s · %s · %s · 浏览%d · 评论%d · %s", author, date, votes, m.question.ViewCount, m.question.CommentCount, statusLabel(m.question.Status))
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
