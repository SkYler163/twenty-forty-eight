package model

import (
	"path/filepath"

	"github.com/adrg/xdg"

	"github.com/SkYler163/twenty-forty-eight/internal/entity"
	"github.com/SkYler163/twenty-forty-eight/internal/storage"
)

type GameModel struct {
	storage      *storage.Storage
	grid         entity.Grid
	savePath     string
	gridSize     int
	score        int
	best         int
	width        int
	height       int
	chance2      uint32
	lost         bool
	victory      bool
	totalVictory bool
	moved        bool
	locked       bool
}

func InitModel(best int) *GameModel {
	m := &GameModel{
		score:    0,
		lost:     false,
		grid:     [entity.GridSize][entity.GridSize]uint16{},
		moved:    true,
		best:     best,
		savePath: filepath.Join(xdg.DataHome, "twenty-forty-eight", "save.gob"),
		chance2:  85,
	}
	m.initializeGrid()
	return m
}

func (m *GameModel) initializeGrid() {
	m.fillRandomCell()
	m.fillRandomCell()
	m.moved = false
}
