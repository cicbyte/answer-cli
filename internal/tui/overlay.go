package tui

import (
	"strings"
)

func overlayError(base string, err error) string {
	msg := styleError.Render("错误: ") + err.Error()
	box := styleErrorBox.Render(msg)
	boxLines := strings.Split(box, "\n")
	w := 0
	for _, l := range boxLines {
		if len(l) > w {
			w = len(l)
		}
	}
	return overlayCenter(base, box, w)
}

func overlayCenter(base, overlay string, w int) string {
	baseLines := strings.Split(base, "\n")
	overlayLines := strings.Split(overlay, "\n")
	maxW := 0
	for _, l := range baseLines {
		if len(l) > maxW {
			maxW = len(l)
		}
	}

	startY := (len(baseLines) - len(overlayLines)) / 2
	if startY < 0 {
		startY = 0
	}
	startX := (maxW - w) / 2
	if startX < 0 {
		startX = 0
	}

	result := make([]string, len(baseLines))
	copy(result, baseLines)
	for i, line := range overlayLines {
		y := startY + i
		if y >= len(result) {
			break
		}
		pad := strings.Repeat(" ", startX)
		result[y] = result[y] + pad + line
	}
	return strings.Join(result, "\n")
}
