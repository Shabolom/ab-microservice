package inMemmoryCashe

import "ab/internal/dto"

func (r *RawExperimentSessionStorage) GetExperimentByNamespace(namespace string) dto.NameSpaceExperiments {
	r.mu.RLock()
	defer r.mu.RUnlock()

	experiments := r.session[namespace]

	return experiments
}
