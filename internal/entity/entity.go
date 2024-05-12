package entity

const (
	GridSize = 4
)

type Grid [GridSize][GridSize]uint16

type SaveState struct {
	Score        int
	Best         int
	Grid         Grid
	Lost         bool
	TotalVictory bool
	Victory      bool
}
