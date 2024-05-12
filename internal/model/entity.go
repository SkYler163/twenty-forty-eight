package model

import (
	"github.com/charmbracelet/lipgloss"
	"os"
)

type direction int8

const (
	directionUp direction = iota
	directionRight
	directionDown
	directionLeft
)

var (
	color0          = lipgloss.AdaptiveColor{Light: "#CCC0B3", Dark: "#E0D7CD"}
	color2          = lipgloss.AdaptiveColor{Light: "#EEE4DA", Dark: "#FFFFFF"}
	color4          = lipgloss.AdaptiveColor{Light: "#EDE0C8", Dark: "#FFFFFF"}
	color8          = lipgloss.AdaptiveColor{Light: "#F2B179", Dark: "#FFD7A9"}
	color16         = lipgloss.AdaptiveColor{Light: "#F59563", Dark: "#FFA875"}
	color32         = lipgloss.AdaptiveColor{Light: "#F67C5F", Dark: "#FF8967"}
	color64         = lipgloss.AdaptiveColor{Light: "#F65E3B", Dark: "#FF7A54"}
	color128        = lipgloss.AdaptiveColor{Light: "#EDCF72", Dark: "#FFE894"}
	color256        = lipgloss.AdaptiveColor{Light: "#EDCC61", Dark: "#FFEE80"}
	color512        = lipgloss.AdaptiveColor{Light: "#EDC850", Dark: "#FFEC70"}
	color1024       = lipgloss.AdaptiveColor{Light: "#EDC53F", Dark: "#FFE960"}
	color2048       = lipgloss.AdaptiveColor{Light: "#EDC22E", Dark: "#FFE750"}
	color4096       = lipgloss.AdaptiveColor{Light: "#3E3933", Dark: "#605E58"}
	headerColor     = lipgloss.AdaptiveColor{Light: "#908D89", Dark: "#F9F6F2"}
	headerFontColor = lipgloss.AdaptiveColor{Light: "#F9F6F2", Dark: "#908D89"}
)

var (
	re          = lipgloss.NewRenderer(os.Stdout)
	headerStyle = re.NewStyle()
	titleStyle  = re.NewStyle().Foreground(headerFontColor).Bold(true).MarginRight(10).MarginLeft(1)
	scoreStyle  = re.NewStyle().
			Foreground(headerFontColor).
			BorderForeground(headerColor).
			Align(lipgloss.Center).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder())
	controlsStyle = re.NewStyle().Italic(true).Foreground(headerFontColor)
	cellStyle     = re.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1).Width(6).Align(lipgloss.Center)
)

var (
	gameLetters = []string{"G", "A", "M", "E"}
	overLetters = []string{"O", "V", "E", "R"}
)
