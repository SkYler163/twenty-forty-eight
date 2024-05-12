package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SkYler163/twenty-forty-eight/internal/entity"
)

func TestStorage(t *testing.T) {
	t.Parallel()

	s := InitStorage("save.gob")
	t.Cleanup(func() {
		err := s.Clear()
		require.NoError(t, err)
	})

	saveState := entity.SaveState{
		Score: 100500,
		Best:  500100,
		Grid: entity.Grid{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
	}

	s.Save(saveState)

	loadedState, err := s.Load()
	require.NoError(t, err)

	assert.EqualValues(t, saveState, loadedState)
}
