package inMemmoryCashe

import (
	"ab/internal/dto"
	"sync"
)

type RawExperimentSessionStorage struct {
	mu      sync.RWMutex
	session map[string]dto.NameSpaceExperiments
}

func New() *RawExperimentSessionStorage {
	return &RawExperimentSessionStorage{
		session: map[string]dto.NameSpaceExperiments{},
	}
}
