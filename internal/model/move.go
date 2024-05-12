package model

import "github.com/SkYler163/twenty-forty-eight/internal/entity"

func (m *GameModel) moveDirection(row, col int, d direction) {
	merged := make(map[int]bool)

	start, end, step := setupDirection(d)

	for pos := start; pos != end; pos += step {
		switch d {
		case directionUp, directionDown:
			row = pos
		case directionLeft, directionRight:
			col = pos
		}

		m.moveCells(row, col, d, merged)
	}
}

func setupDirection(d direction) (int, int, int) {
	var start, end, step int
	switch d {
	case directionUp, directionLeft:
		start, end, step = 1, entity.GridSize, 1
	case directionRight, directionDown:
		start, end, step = entity.GridSize-2, -1, -1
	}

	return start, end, step
}

func (m *GameModel) moveCells(row, col int, d direction, merged map[int]bool) {
	switch d {
	case directionUp:
		for i := row - 1; i >= 0; i-- {
			if !m.tryMergeOrMove(row, col, i, col, merged) {
				break
			}
			row = i // Обновляем текущую позицию тайла
		}
	case directionDown:
		for i := row + 1; i < entity.GridSize; i++ {
			if !m.tryMergeOrMove(row, col, i, col, merged) {
				break
			}
			row = i
		}
	case directionLeft:
		for i := col - 1; i >= 0; i-- {
			if !m.tryMergeOrMove(row, col, row, i, merged) {
				break
			}
			col = i
		}
	case directionRight:
		for i := col + 1; i < entity.GridSize; i++ {
			if !m.tryMergeOrMove(row, col, row, i, merged) {
				break
			}
			col = i
		}
	}
}

func (m *GameModel) tryMergeOrMove(fromRow, fromCol, toRow, toCol int, merged map[int]bool) bool {
	switch m.grid[toRow][toCol] {
	case 0:
		m.grid[toRow][toCol], m.grid[fromRow][fromCol] = m.grid[fromRow][fromCol], 0
		m.moved = true
		return true
	case m.grid[fromRow][fromCol]:
		if merged[toRow*entity.GridSize+toCol] || merged[fromRow*entity.GridSize+fromCol] {
			return false
		}

		m.grid[toRow][toCol] *= 2
		m.grid[fromRow][fromCol] = 0
		m.score += int(m.grid[toRow][toCol])
		m.moved = true

		return false
	}

	return false
}

func (m *GameModel) canMoveUp() bool {
	for col := 0; col < entity.GridSize; col++ {
		for row := 1; row < entity.GridSize; row++ {
			if m.grid[row][col] == 0 {
				continue
			}

			if m.grid[row-1][col] == 0 {
				return true
			}

			if m.grid[row-1][col] == m.grid[row][col] {
				return true
			}
		}
	}

	return false
}

func (m *GameModel) canMoveDown() bool {
	for col := 0; col < entity.GridSize; col++ {
		for row := entity.GridSize - 2; row >= 0; row-- {
			if m.grid[row][col] == 0 {
				continue
			}

			if m.grid[row+1][col] == 0 {
				return true
			}

			if m.grid[row+1][col] == m.grid[row][col] {
				return true
			}
		}
	}

	return false
}

func (m *GameModel) canMoveRight() bool {
	for row := 0; row < entity.GridSize; row++ {
		for col := entity.GridSize - 2; col >= 0; col-- {
			if m.grid[row][col] == 0 {
				continue
			}

			if m.grid[row][col+1] == 0 {
				return true
			}

			if m.grid[row][col+1] == m.grid[row][col] {
				return true
			}
		}
	}

	return false
}

func (m *GameModel) canMoveLeft() bool {
	for row := 0; row < entity.GridSize; row++ {
		for col := 1; col < entity.GridSize; col++ {
			if m.grid[row][col] == 0 {
				continue
			}

			if m.grid[row][col-1] == 0 {
				return true
			}

			if m.grid[row][col-1] == m.grid[row][col] {
				return true
			}
		}
	}

	return false
}
