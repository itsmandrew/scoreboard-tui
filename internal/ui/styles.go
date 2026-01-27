package ui

import "github.com/charmbracelet/lipgloss"

var (
	MainColor   = lipgloss.Color("#7D56F4")
	WindowStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(MainColor).
			Padding(1, 8).
			Margin(2)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(MainColor).
			Padding(0, 1).
			Bold(true)
)
