package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/models"
)

func RunTUI() error {
	cli, err := NewTUIClient()
	if err != nil {
		return err
	}
	p := tea.NewProgram(NewAppModel(cli), tea.WithAltScreen(), tea.WithMouseAllMotion())
	_, err = p.Run()
	return err
}

type appModel struct {
	cli        *client.Client
	width      int
	height     int
	activeView viewKind

	// List state
	questions    []models.QuestionListItem
	qCount       int64
	qPage        int
	qPageSize     int
	qCursor      int
	qLoading     bool
	qOrder       string
	searchMode   bool
	searchInput  textinput.Model
	searchQuery  string
	searchActive bool

	// Detail state
	question          *models.QuestionInfoResp
	answers           []models.AnswerInfo
	comments          map[string][]models.CommentInfo
	dLoading          bool
	viewport          viewport.Model
	detailContent     string
	detailHeaderLines int

	// Editor state
	editorActive bool
	editorMode   editorMode
	editorInput  textarea.Model
	editorTarget string

	// Auth
	username string

	// Error overlay
	err error
}

func NewAppModel(cli *client.Client) appModel {
	si := textinput.New()
	si.Prompt = "/"
	si.PromptStyle = styleCursor
	si.CharLimit = 200
	si.Width = 40

	return appModel{
		cli:         cli,
		qPageSize:    20,
		qOrder:      "newest",
		comments:    make(map[string][]models.CommentInfo),
		searchInput: si,
		viewport:    viewport.New(0, 0),
	}
}

func (m appModel) Init() tea.Cmd {
	return tea.Batch(checkAuthCmd(m.cli), loadQuestionsCmd(m.cli, 1, m.qPageSize, m.qOrder))
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Global handlers
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width
		if m.detailHeaderLines > 0 {
			m.viewport.Height = m.height - m.detailHeaderLines - 3
		} else {
			m.viewport.Height = m.height
		}
		return m, nil

	case tea.KeyMsg:
		if m.editorActive {
			return m.updateEditor(msg)
		}
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.err != nil {
			if msg.String() == "enter" || msg.String() == "esc" {
				m.err = nil
				return m, nil
			}
			return m, nil
		}

	case apiErrorMsg:
		m.dLoading = false
		m.qLoading = false
		if client.IsUnauthorizedError(msg.err) {
			m.err = fmt.Errorf("未登录或登录已过期，请先执行: answer-cli auth login")
			return m, nil
		}
		m.err = msg.err
		return m, nil

	case authCheckedMsg:
		m.username = msg.username
		return m, nil

	case voteResultMsg:
		if msg.err != nil {
			m.err = msg.err
		}
		return m, nil
	}

	// View-specific handlers
	if m.activeView == viewQuestionDetail {
		return m.updateQuestionDetail(msg)
	}

	// List view handles remaining messages
	return m.updateQuestionList(msg)
}

func (m appModel) View() string {
	var content string
	switch m.activeView {
	case viewQuestionList:
		content = m.viewQuestionList()
	case viewQuestionDetail:
		content = m.viewQuestionDetail()
	}

	content += "\n" + m.renderStatusBar()

	if m.editorActive {
		content = overlayEditor(content, m)
	}
	if m.err != nil {
		content = overlayError(content, m.err)
	}
	return content
}

func (m appModel) renderStatusBar() string {
	left := "未登录"
	if m.username != "" {
		left = styleSuccess.Render(m.username)
	}

	center := ""
	switch m.activeView {
	case viewQuestionList:
		if m.searchActive && m.searchQuery != "" {
			center = fmt.Sprintf("搜索: %s", m.searchQuery)
		} else {
			center = fmt.Sprintf("问题列表 (%d) | 排序: %s", m.qCount, m.orderLabel())
		}
	case viewQuestionDetail:
		center = "问题详情"
	}

	right := ""
	switch m.activeView {
	case viewQuestionList:
		if m.searchMode {
			right = "Enter:搜索 Esc:取消"
		} else {
			right = "↑↓:移动 Enter:查看 /:搜索 Tab:排序 n/p:翻页 q:退出"
		}
	case viewQuestionDetail:
		right = "↑↓:滚动 Esc:返回 v:投票 r:回答 c:评论 g:顶部 G:底部 q:退出"
	}

	barW := m.width
	if barW <= 0 {
		barW = 80
	}

	leftW := lipgloss.Width(left)
	centerW := lipgloss.Width(center)
	rightW := lipgloss.Width(right)
	used := leftW + centerW + rightW
	if used > barW {
		return styleStatusBar.Render(left)
	}
	gap := max((barW-used)/2, 1)

	return styleStatusBar.Render(left + strings.Repeat(" ", gap) + center + strings.Repeat(" ", gap) + right)
}

func (m appModel) orderLabel() string {
	switch m.qOrder {
	case "newest":
		return "最新"
	case "active":
		return "活跃"
	case "hot":
		return "热门"
	case "score":
		return "评分"
	default:
		return m.qOrder
	}
}
