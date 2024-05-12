package model

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SkYler163/twenty-forty-eight/internal/storage"
)

func (m *GameModel) Init() tea.Cmd {
	m.storage = storage.InitStorage(m.savePath)
	state, err := m.storage.Load()
	if err != nil {
		return nil
	}

	m.grid = state.Grid
	m.score = state.Score
	m.best = state.Best
	m.lost = state.Lost
	m.totalVictory = state.TotalVictory

	return nil
}
