package model

import (
	"math/rand/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SkYler163/twenty-forty-eight/internal/entity"
	"github.com/SkYler163/twenty-forty-eight/internal/storage"
)

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.locked {
		return m, nil
	}

	m.locked = true
	defer func() {
		m.locked = false
	}()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if cmd := m.handleKey(msg); cmd != nil {
			return m, cmd
		}
	}

	if m.lost || m.totalVictory {
		return m, nil
	}

	m.postUpdateChecks()
	m.storage.Save(
		entity.SaveState{Score: m.score, Best: m.best, Grid: m.grid, TotalVictory: m.totalVictory, Lost: m.lost},
	)

	return m, nil
}

func (m *GameModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	switch msg.Type {
	case tea.KeyCtrlC:
		return tea.Quit
	case tea.KeyCtrlR:
		m.resetModel()

		return nil
	case tea.KeyCtrlE:
		m.best = 0
		m.resetModel()

		return nil
	case tea.KeyUp, tea.KeyRight, tea.KeyDown, tea.KeyLeft:
		if m.lost {
			return nil
		}

		switch msg.Type {
		case tea.KeyUp:
			m.moveUp()
		case tea.KeyRight:
			m.moveRight()
		case tea.KeyDown:
			m.moveDown()
		case tea.KeyLeft:
			m.moveLeft()
		}
	}

	return nil
}

func (m *GameModel) resetModel() {
	*m = *InitModel(m.best)
	m.storage = storage.InitStorage(m.savePath)
	m.storage.Save(entity.SaveState{Score: m.score, Best: m.best, Grid: m.grid})
}

func (m *GameModel) postUpdateChecks() {
	if m.moved {
		m.fillRandomCell()
		m.moved = false
	}

	m.checkLost()
	if m.lost {
		return
	}

	if m.score > m.best {
		m.best = m.score
	}
}

func (m *GameModel) moveUp() {
	for c := 0; c < entity.GridSize; c++ {
		m.moveDirection(1, c, directionUp)
	}
}

func (m *GameModel) moveRight() {
	for r := 0; r < entity.GridSize; r++ {
		m.moveDirection(r, 2, directionRight)
	}
}

func (m *GameModel) moveDown() {
	for c := 0; c < entity.GridSize; c++ {
		m.moveDirection(2, c, directionDown)
	}
}

func (m *GameModel) moveLeft() {
	for r := 0; r < entity.GridSize; r++ {
		m.moveDirection(r, 1, directionLeft)
	}
}

func (m *GameModel) checkLost() {
	m.lost = true
	for _, canMove := range []func() bool{m.canMoveUp, m.canMoveRight, m.canMoveDown, m.canMoveLeft} {
		if canMove() {
			m.lost = false
			break
		}
	}
}

// fillRandomCell fills empty cell using reservoir sampling algorithm
func (m *GameModel) fillRandomCell() {
	if !m.moved || m.isGridFull() {
		return
	}

	var chosenR, chosenC int
	count := 0

	rnd := rand.New(rand.NewPCG(0, uint64(time.Now().UnixNano())))

	for r := 0; r < entity.GridSize; r++ {
		for c := 0; c < entity.GridSize; c++ {
			if m.grid[r][c] == 0 {
				count++

				if rnd.IntN(count) == 0 {
					chosenR, chosenC = r, c
				}
			}
		}
	}

	if count == 0 {
		return
	}

	if rnd.Uint32N(100) > m.chance2 {
		m.grid[chosenR][chosenC] = 4

		return
	}

	m.grid[chosenR][chosenC] = 2
}

func (m *GameModel) isGridFull() bool {
	for r := 0; r < entity.GridSize; r++ {
		for c := 0; c < entity.GridSize; c++ {
			if m.grid[r][c] == 0 {
				return false
			}
		}
	}

	return true
}
