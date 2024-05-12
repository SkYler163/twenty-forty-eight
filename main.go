package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SkYler163/twenty-forty-eight/internal/model"
)

func main() {
	p := tea.NewProgram(model.InitModel(0), tea.WithAltScreen(), tea.WithOutput(os.Stderr))

	if _, err := p.Run(); err != nil {
		log.Println("failed to run program:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
