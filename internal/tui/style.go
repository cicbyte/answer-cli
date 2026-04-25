package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorPrimary   = lipgloss.Color("#61AFEF")
	colorSecondary = lipgloss.Color("#C678DD")
	colorSuccess   = lipgloss.Color("#98C379")
	colorWarning   = lipgloss.Color("#E5C07B")
	colorError     = lipgloss.Color("#E06C75")
	colorMuted     = lipgloss.Color("#5C6370")

	styleTitle = lipgloss.NewStyle().Bold(true).Foreground(colorPrimary)
	styleCursor = lipgloss.NewStyle().Foreground(colorSecondary).Bold(true)
	styleMuted = lipgloss.NewStyle().Foreground(colorMuted)
	styleSuccess = lipgloss.NewStyle().Foreground(colorSuccess)
	styleError = lipgloss.NewStyle().Foreground(colorError)
	styleTag = lipgloss.NewStyle().Foreground(colorPrimary).Padding(0, 1)
	styleVote = lipgloss.NewStyle().Foreground(colorWarning).Bold(true)
	styleAccepted = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true)
	styleDim = lipgloss.NewStyle().Foreground(lipgloss.Color("#4B5263"))

	styleStatusBar = lipgloss.NewStyle().
		Background(lipgloss.Color("#21252B")).
		Foreground(colorMuted).
		Padding(0, 1)

	styleErrorBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorError).
		Padding(1, 2)

	styleSeparator = lipgloss.NewStyle().Foreground(colorMuted)
)

func maxWidth(w int) int {
	if w <= 4 {
		return 80
	}
	return w - 4
}
