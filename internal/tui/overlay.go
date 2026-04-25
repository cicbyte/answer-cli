package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func overlayError(base string, err error) string {
	msg := styleError.Render("错误: ") + err.Error() + "\n" + styleMuted.Render("Enter/Esc 关闭  Ctrl+C 退出")
	box := styleErrorBox.Render(msg)
	boxLines := strings.Split(box, "\n")
	w := 0
	for _, l := range boxLines {
		lw := lipgloss.Width(l)
		if lw > w {
			w = lw
		}
	}
	return overlayCenter(base, box, w)
}

func overlayCenter(base, overlay string, w int) string {
	baseLines := strings.Split(base, "\n")
	overlayLines := strings.Split(overlay, "\n")

	termW := 80
	for _, l := range baseLines {
		lw := lipgloss.Width(l)
		if lw > termW {
			termW = lw
		}

	}

	startY := (len(baseLines) - len(overlayLines)) / 2
	if startY < 0 {
		startY = 0
	}
	startX := max((termW-w)/2, 0)

	pad := strings.Repeat(" ", startX)
	result := make([]string, len(baseLines))
	for i, line := range baseLines {
		lw := lipgloss.Width(line)
		result[i] = line + strings.Repeat(" ", max(termW-lw, 0))
	}

	for i, line := range overlayLines {
		y := startY + i
		if y >= len(result) {
			break
		}
		result[y] = pad + line + strings.Repeat(" ", max(termW-startX-lipgloss.Width(line), 0))
	}

	return strings.Join(result, "\n")
}
