package tui

import (
	"github.com/charmbracelet/glamour"
)

func renderMarkdown(content string, width int) string {
	if content == "" {
		return ""
	}
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return content
	}
	out, err := r.Render(content)
	if err != nil {
		return content
	}
	return out
}
