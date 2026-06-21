package inMemmoryCashe

import "ab/internal/dto"

func (r *RawExperimentSessionStorage) ReplaceFeatureToggle(newCash map[string]dto.NamespaceFeatureToggle) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sessionFeatureToggle = newCash
}
