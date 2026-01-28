package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/itsmandrew/scoreboard-cli/internal/sports"
)

type sessionState int

const (
	menuView sessionState = iota
	resultView
)

type Model struct {
	state    sessionState
	choices  []string
	cursor   int
	selected string
	loading  bool
	spinner  spinner.Model
	apiKey   string
	nbaGames []sports.Game
	err      error
}

func InitialModel(apiKey string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(MainColor)

	return Model{
		state:   menuView,
		choices: []string{"NFL Games", "NBA Games", "NCAA Basketball", "Exit"},
		spinner: s,
		apiKey:  apiKey,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		if m.loading {
			return m, nil
		}

		return m.handleKeyInput(msg)

	case string:
		m.loading = false
		m.state = resultView
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case nbaMsg:
		m.loading = false
		m.state = resultView
		m.nbaGames = msg
		return m, nil

	case errMsg:
		m.loading = false
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m *Model) handleKeyInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.state == resultView {
		if msg.String() == "enter" {
			m.state = menuView
		}
		return m, nil
	}

	// Menu Navigation Logic
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		m.selected = m.choices[m.cursor]
		if m.selected == "Exit" {
			return m, tea.Quit
		}

		m.loading = true
		if m.selected == "NBA Games" {
			return m, tea.Batch(m.spinner.Tick, FetchNBACmd(m.apiKey))
		}
		return m, tea.Batch(m.spinner.Tick, m.simulateFetch())
	}

	return m, nil
}

func (m Model) simulateFetch() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return "Done"
	}
}

func (m Model) renderNBAGames() string {
	if len(m.nbaGames) == 0 {
		return "No games scheduled for today"
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("NBA Scores - %s\n\n", time.Now().Format("Jan 02, 2006")))

	for _, game := range m.nbaGames {

		homeStr := fmt.Sprintf("%s %d", game.HomeTeam.Abbreviation, game.HomeTeamScore)
		visitorStr := fmt.Sprintf("%s %d", game.VisitorTeam.Abbreviation, game.VisitorTeamScore)

		status := sports.FormatStatus(game.Status)

		row := fmt.Sprintf("%s vs %s — %s\n",
			lipgloss.NewStyle().Width(10).Render(homeStr),
			lipgloss.NewStyle().Width(10).Render(visitorStr),
			lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(status),
		)
		b.WriteString(row)
	}
	return b.String()
}

func (m Model) View() string {
	var body string
	if m.loading {
		body = fmt.Sprintf(
			"\n %s  Fetching %s data...\n\n%s",
			m.spinner.View(),
			m.selected,
			lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("Please wait..."),
		)
	} else if m.state == resultView {
		// The "Message" Screen
		content := ""
		if m.selected == "NBA Games" {
			content = m.renderNBAGames()
		} else {
			content = fmt.Sprintf("Returning data for: %s", m.selected)
		}

		body = fmt.Sprintf(
			"%s\n\n%s\n\n%s",
			TitleStyle.Render(" SCOREBOARD "),
			content,
			lipgloss.NewStyle().Italic(true).Render("Press Enter to go back..."),
		)

	} else {
		// The "Menu" Screen
		var b strings.Builder
		b.WriteString(TitleStyle.Render(" SCOREBOARD CLI MENU ") + "\n\n")

		for i, choice := range m.choices {
			cursor := " "
			label := choice

			if m.cursor == i {
				cursor = ">"
				label = lipgloss.NewStyle().Foreground(MainColor).Bold(true).Render(choice)
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, label))
		}

		b.WriteString("\n(j/k to move • enter to select)")
		body = b.String()
	}

	return WindowStyle.Render(body)
}
