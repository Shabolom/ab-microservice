package inMemmoryCashe

import "ab/internal/dto"

func (r *RawExperimentSessionStorage) ReplaceWithoutCustomGroups(newCash map[string]dto.NameSpaceExperiments) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sessionWithoutCustomParamGroups = newCash
}
