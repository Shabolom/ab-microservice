package inMemmoryCashe

import (
	"ab/internal/dto"
)

func (r *RawExperimentSessionStorage) Replace(newCash map[string]dto.NameSpaceExperiments) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.session = newCash
}
