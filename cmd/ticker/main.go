package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/itsmandrew/scoreboard-cli/internal/ui"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's an error: %v", err)
		os.Exit(1)
	}
}
