package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textarea"
)

func newEditorTextarea(placeholder string) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = placeholder
	ta.CharLimit = 5000
	ta.SetWidth(60)
	ta.SetHeight(8)
	ta.ShowLineNumbers = false
	return ta
}

func (m appModel) updateEditor(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			content := m.editorInput.Value()
			if content == "" {
				return m, nil
			}
			switch m.editorMode {
			case modeAnswer:
				return m, addAnswerCmd(m.cli, m.editorTarget, content)
			case modeComment:
				return m, addCommentCmd(m.cli, m.editorTarget, content)
			}
		case "esc":
			m.editorActive = false
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.editorInput, cmd = m.editorInput.Update(msg)
	return m, cmd
}

func overlayEditor(base string, m appModel) string {
	var title string
	switch m.editorMode {
	case modeAnswer:
		title = "编写回答"
	case modeComment:
		title = "添加评论"
	}

	box := styleEditor.Render(
		styleTitle.Render(title) + "\n\n" +
			m.editorInput.View() + "\n\n" +
			styleMuted.Render("Ctrl+S:提交  Esc:取消"),
	)

	w := len(box)
	if w <= 0 {
		w = 60
	}
	return overlayCenter(base, box, w)
}
