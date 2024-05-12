package model

import (
	"github.com/SkYler163/twenty-forty-eight/internal/entity"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func (m *GameModel) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.header(),
			m.rows(),
			controlsStyle.Render(
				"Controls:\n"+
					"- Arrow Keys: Move tiles in the specified direction.\n"+
					"- Ctrl+C: Exit the game.\n"+
					"- Ctrl+R: Restart the game.\n"+
					"- Ctrl+E: Reset the best score and restart the game.\n\n"+
					"Author: \u001B]8;;https://www.linkedin.com/in/sergey-vertepov\aSergey Vertepov\u001B]8;;\a\u001B[0m",
			),
		),
	)
}

func (m *GameModel) header() string {
	return headerStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			titleStyle.Render("2048"),
			scoreContainer("SCORE", m.score),
			scoreContainer("BEST", m.best),
		))
}

func scoreContainer(title string, val int) string {
	return scoreStyle.
		Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				lipgloss.NewStyle().Align(lipgloss.Center).Render(title),
				lipgloss.NewStyle().Align(lipgloss.Center).Render(strconv.Itoa(val)),
			),
		)
}

func (m *GameModel) rows() string {
	if m.lost {
		return m.renderLoss()
	}

	return m.renderGrid()
}

func (m *GameModel) renderGrid() string {
	rows := make([]string, 0, entity.GridSize)
	for r := 0; r < entity.GridSize; r++ {
		rowCells := make([]string, 0, entity.GridSize)
		for c := 0; c < entity.GridSize; c++ {
			val := ""
			if m.grid[r][c] > 0 {
				val = strconv.Itoa(int(m.grid[r][c]))
			}
			rowCells = append(rowCells, cellView(colorByValue(m.grid[r][c]), val))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rowCells...))
	}

	return lipgloss.JoinVertical(lipgloss.Top, rows...)
}

func (m *GameModel) renderLoss() string {
	gameRow, overRow := entity.GridSize/2-1, entity.GridSize/2
	startCol := (entity.GridSize - 4) / 2
	rows := make([]string, entity.GridSize)

	for r := 0; r < entity.GridSize; r++ {
		var rowCells []string
		for c := 0; c < entity.GridSize; c++ {
			val := ""
			color := colorByValue(m.grid[r][c])
			if r == gameRow || r == overRow {
				if c >= startCol && c < startCol+4 {
					color = color64
					if r == gameRow {
						val = gameLetters[c-startCol]
					} else {
						val = overLetters[c-startCol]
					}
				}
			} else {
				if m.grid[r][c] > 0 {
					val = strconv.Itoa(int(m.grid[r][c]))
				}
			}
			rowCells = append(rowCells, cellView(color, val))
		}
		rows[r] = lipgloss.JoinHorizontal(lipgloss.Top, rowCells...)
	}

	return lipgloss.JoinVertical(lipgloss.Top, rows...)
}

func cellView(color lipgloss.AdaptiveColor, val string) string {
	return cellStyle.Foreground(color).BorderForeground(color).Render(val)
}

func colorByValue(value uint16) lipgloss.AdaptiveColor {
	var (
		color lipgloss.AdaptiveColor
	)

	switch value {
	case 0:
		color = color0
	case 2:
		color = color2
	case 4:
		color = color4
	case 8:
		color = color8
	case 16:
		color = color16
	case 32:
		color = color32
	case 64:
		color = color64
	case 128:
		color = color128
	case 256:
		color = color256
	case 512:
		color = color512
	case 1024:
		color = color1024
	case 2048:
		color = color2048
	case 4096:
		color = color4096
	default:
		color = color0
	}

	return color
}
