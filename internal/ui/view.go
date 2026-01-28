package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-cli/internal/sports"
)

func (m Model) View() string {
	var body string

	// Route rendering based on application state
	switch {
	case m.err != nil:
		body = m.errorView()
	case m.loading:
		body = m.loadingView()
	case m.state == resultView:
		body = m.resultView()
	default:
		body = m.menuView()
	}

	return WindowStyle.Render(body)
}

func (m Model) menuView() string {
	var b strings.Builder
	b.WriteString(TitleStyle.Render(" SCOREBOARD CLI MENU ") + "\n\n")

	for i, choice := range m.choices {
		cursor := "  "
		label := choice

		if m.cursor == i {
			cursor = "> "
			label = lipgloss.NewStyle().Foreground(MainColor).Bold(true).Render(choice)
		}
		b.WriteString(fmt.Sprintf("%s%s\n", cursor, label))
	}

	b.WriteString(fmt.Sprintf("\n%s", SubtleStyle.Render("(j/k to move • enter to select)")))
	return b.String()
}

func (m Model) loadingView() string {
	return fmt.Sprintf(
		"\n %s  Fetching %s data...\n\n%s",
		m.spinner.View(),
		m.selected,
		SubtleStyle.Render("Please wait..."),
	)
}

func (m Model) resultView() string {
	content := ""
	if m.selected == "NBA Games" {
		content = m.renderNBAGames()
	} else {
		content = fmt.Sprintf("Returning data for: %s", m.selected)
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		TitleStyle.Render(" SCOREBOARD "),
		content,
		ItalicStyle.Render("Press Enter to go back..."),
	)
}

func (m Model) errorView() string {
	return fmt.Sprintf(
		"%s\n\nError: %v\n\n%s",
		ErrorTitleStyle.Render(" ERROR "),
		m.err,
		ItalicStyle.Render("Press Enter to go back..."),
	)
}

// Format fetched NBA game data into a table-like string
func (m Model) renderNBAGames() string {
	if len(m.nbaGames) == 0 {
		return SubtleStyle.Render("No games scheduled for today")
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("NBA Scores - %s\n\n", time.Now().Format("Jan 02, 2006")))

	for _, game := range m.nbaGames {
		home := TeamStyle.Render(fmt.Sprintf("%s %d", game.HomeTeam.Abbreviation, game.HomeTeamScore))
		away := TeamStyle.Render(fmt.Sprintf("%s %d", game.VisitorTeam.Abbreviation, game.VisitorTeamScore))
		status := StatusStyle.Render(sports.FormatStatus(game.Status))

		b.WriteString(fmt.Sprintf("%s vs %s — %s\n", home, away, status))
	}
	return b.String()
}
