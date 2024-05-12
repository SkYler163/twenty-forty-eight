package storage

import (
	"encoding/gob"
	"os"
	"path/filepath"

	"github.com/SkYler163/twenty-forty-eight/internal/entity"
)

type Storage struct {
	path string
}

func InitStorage(path string) *Storage {
	return &Storage{path: path}
}

func (s *Storage) Save(save entity.SaveState) {
	dir := filepath.Dir(s.path)
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0o755); err != nil {
			return
		}
	}

	file, err := os.Create(s.path)
	if err != nil {
		return
	}
	defer file.Close()

	if err := gob.NewEncoder(file).Encode(save); err != nil {
		return
	}
}

func (s *Storage) Load() (entity.SaveState, error) {
	file, err := os.Open(s.path)
	if err != nil {
		return entity.SaveState{}, err
	}
	defer file.Close()

	var save entity.SaveState
	if err := gob.NewDecoder(file).Decode(&save); err != nil {
		return entity.SaveState{}, err
	}

	return save, nil
}

func (s *Storage) Clear() error {
	return os.Remove(s.path)
}
