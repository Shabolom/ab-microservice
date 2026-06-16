package inMemmoryCashe

import (
	"ab/internal/dto"
	"sync"
)

type RawExperimentSessionStorage struct {
	mu                              sync.RWMutex
	sessionWithCustomParamGroups    map[string]dto.NameSpaceExperiments
	sessionWithoutCustomParamGroups map[string]dto.NameSpaceExperiments
}

func New() *RawExperimentSessionStorage {
	return &RawExperimentSessionStorage{
		sessionWithCustomParamGroups:    map[string]dto.NameSpaceExperiments{},
		sessionWithoutCustomParamGroups: map[string]dto.NameSpaceExperiments{},
	}
}
