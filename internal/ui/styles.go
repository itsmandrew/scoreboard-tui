package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	MainColor   = lipgloss.Color("#7D56F4")
	SubtleColor = lipgloss.Color("241")
	ErrorColor  = lipgloss.Color("#FF5555")
	White       = lipgloss.Color("#FFFDF5")

	// Base Containers
	WindowStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(MainColor).
			Padding(1, 4).
			Margin(2)

	// Text Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(White).
			Background(MainColor).
			Padding(0, 1).
			Bold(true)

	ErrorTitleStyle = TitleStyle.Copy().
			Background(ErrorColor)

	// Added missing subtle and italic styles
	SubtleStyle = lipgloss.NewStyle().
			Foreground(SubtleColor)

	ItalicStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(SubtleColor)

	// Scoreboard Layout Styles
	TeamStyle = lipgloss.NewStyle().
			Width(12).
			Bold(true)

	StatusStyle = lipgloss.NewStyle().
			Foreground(SubtleColor).
			MarginLeft(2)
)
