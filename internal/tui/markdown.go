package tui

import (
	"github.com/cicbyte/answer-cli/internal/output"
)

func renderMarkdown(content string, width int) string {
	return output.RenderMarkdownWidth(content, width)
}
