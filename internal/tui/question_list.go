package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cicbyte/answer-cli/internal/models"
)

func (m appModel) updateQuestionList(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.searchMode {
		var cmd tea.Cmd
		m.searchInput, cmd = m.searchInput.Update(msg)
		if cmd != nil {
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.searchMode {
			switch msg.String() {
			case "enter":
				q := strings.TrimSpace(m.searchInput.Value())
				if q != "" {
					m.qLoading = true
					m.searchActive = false
					return m, searchCmd(m.cli, q)
				}
			case "esc":
				m.searchMode = false
				m.searchInput.SetValue("")
			}
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.qCursor > 0 {
				m.qCursor--
			}
		case "down", "j":
			if m.qCursor < len(m.questions)-1 {
				m.qCursor++
			}
		case "enter":
			if m.qCursor < len(m.questions) {
				q := m.questions[m.qCursor]
				m.activeView = viewQuestionDetail
				m.dLoading = true
				m.answers = nil
				m.comments = make(map[string][]models.CommentInfo)
				return m, loadQuestionDetailCmd(m.cli, q.ID)
			}
		case "n":
			if !m.searchActive {
				totalPages := int(m.qCount) / m.qPageSize
				if int(m.qCount)%m.qPageSize > 0 {
					totalPages++
				}
				if m.qPage < totalPages {
					m.qLoading = true
					m.qPage++
					return m, loadQuestionsCmd(m.cli, m.qPage, m.qPageSize, m.qOrder)
				}
			}
		case "p":
			if !m.searchActive && m.qPage > 1 {
				m.qLoading = true
				m.qPage--
				return m, loadQuestionsCmd(m.cli, m.qPage, m.qPageSize, m.qOrder)
			}
		case "/":
			m.searchMode = true
			m.searchInput.SetValue("")
			return m, m.searchInput.Focus()
		case "tab":
			orders := []string{"newest", "active", "hot", "score"}
			for i, o := range orders {
				if o == m.qOrder {
					m.qOrder = orders[(i+1)%len(orders)]
					break
				}
			}
			m.searchActive = false
			m.qLoading = true
			m.qPage = 1
			return m, loadQuestionsCmd(m.cli, 1, m.qPageSize, m.qOrder)
		}
	}
	return m, nil
}

func (m appModel) viewQuestionList() string {
	var lines []string

	header := styleTitle.Render("问题列表")
	if m.searchActive && m.searchQuery != "" {
		header = styleTitle.Render(fmt.Sprintf("搜索: %s", m.searchQuery))
	}
	if !m.searchActive {
		header += styleMuted.Render(fmt.Sprintf("  (第 %d 页, 共 %d 条)", m.qPage, m.qCount))
	}
	lines = append(lines, header)

	if m.searchMode {
		lines = append(lines, m.searchInput.View())
	}

	if m.qLoading {
		lines = append(lines, "", styleMuted.Render("加载中..."))
		return strings.Join(lines, "\n")
	}

	if len(m.questions) == 0 {
		lines = append(lines, "", styleMuted.Render("没有找到问题"))
		return strings.Join(lines, "\n")
	}

	availW := m.width
	if availW <= 0 {
		availW = 80
	}
	availW -= 4

	for i, q := range m.questions {
		cursor := "  "
		if i == m.qCursor {
			cursor = styleCursor.Render("> ")
		}

		title := q.Title
		titleRunes := []rune(title)
		if len(titleRunes) > availW-30 {
			title = string(titleRunes[:availW-30]) + "..."
		}

		date := models.FormatTimestamp(q.CreatedAt).Format("01-02")
		author := ""
		if q.UserInfo != nil {
			author = q.UserInfo.DisplayName
		}

		vote := styleVote.Render(fmt.Sprintf("%d", q.VoteCount))
		ansCount := fmt.Sprintf("%d", q.AnswerCount)
		tags := ""
		for _, t := range q.Tags {
			tags += styleTag.Render(t.DisplayName)
		}

		line := fmt.Sprintf("%s%s %s  答案:%s  %s %s %s",
			cursor, title, vote, ansCount, tags, date, styleDim.Render(author))

		lineRunes := []rune(line)
		if len(lineRunes) > availW {
			line = string(lineRunes[:availW]) + "..."
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
