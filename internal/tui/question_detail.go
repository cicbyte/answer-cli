package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
	content := m.renderDetailContent()
	if content == "" {
		content = styleMuted.Render("无内容")
	}
	if m.detailScroll > 0 {
		lines := strings.Split(content, "\n")
		if m.detailScroll >= len(lines) {
			m.detailScroll = len(lines) - 1
		}
		content = strings.Join(lines[m.detailScroll:], "\n")
	}
	return content
}

func (m appModel) renderDetailContent() string {
	if m.question == nil {
		return ""
	}

	var b strings.Builder
	w := maxWidth(m.width)

	b.WriteString(styleTitle.Render(m.question.Title))
	b.WriteString("\n\n")

	author := "匿名"
	if m.question.UserInfo != nil {
		author = m.question.UserInfo.DisplayName
	}
	date := models.FormatTimestamp(m.question.CreatedAt).Format("2006-01-02 15:04")
	votes := styleVote.Render(fmt.Sprintf("↑ %d", m.question.VoteCount))
	views := fmt.Sprintf("浏览 %d", m.question.ViewCount)

	meta := fmt.Sprintf("%s  %s  %s", author, date, views)
	b.WriteString(styleMuted.Render(meta))
	b.WriteString("  ")
	b.WriteString(votes)

	if len(m.question.Tags) > 0 {
		b.WriteString("\n")
		for _, t := range m.question.Tags {
			b.WriteString(styleTag.Render(t.DisplayName))
		}
	}
	b.WriteString("\n\n")

	b.WriteString(renderMarkdown(m.question.Content, w))
	b.WriteString("\n")

	b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
	b.WriteString("\n")

	if len(m.answers) == 0 {
		b.WriteString(styleMuted.Render("暂无回答"))
	} else {
		b.WriteString(fmt.Sprintf("回答 (%d)\n", len(m.answers)))
	}

	acceptedMark := styleAccepted.Render(" ✓ 已采纳")

	for i, a := range m.answers {
		if i > 0 {
			b.WriteString("\n")
			b.WriteString(styleSeparator.Render(strings.Repeat("─", w)))
			b.WriteString("\n")
		}

		aAuthor := "匿名"
		if a.UserInfo != nil {
			aAuthor = a.UserInfo.DisplayName
		}
		aDate := models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04")
		aVotes := styleVote.Render(fmt.Sprintf("↑ %d", a.VoteCount))

		header := fmt.Sprintf("回答 %d  %s  %s  %s", i+1, aAuthor, aDate, aVotes)
		if a.ID == m.question.AcceptedAnswerID {
			header += acceptedMark
		}
		b.WriteString(header)
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

	qComments := m.comments[m.question.ID]
	if len(qComments) > 0 {
		b.WriteString("\n")
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
