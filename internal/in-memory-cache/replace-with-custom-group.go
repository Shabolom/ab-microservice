package inMemmoryCashe

import (
	"ab/internal/dto"
)

func (r *RawExperimentSessionStorage) ReplaceWithCustomGroups(newCash map[string]dto.NameSpaceExperiments) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sessionWithCustomParamGroups = newCash
}
