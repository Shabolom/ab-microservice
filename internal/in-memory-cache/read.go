package inMemmoryCashe

import (
	"ab/internal/dto"
)

func (r *RawExperimentSessionStorage) GetExperimentWithCustomGroupsByNamespace(namespace string) dto.NameSpaceExperiments {
	r.mu.RLock()
	defer r.mu.RUnlock()

	experiments := r.sessionWithCustomParamGroups[namespace]

	return experiments
}

func (r *RawExperimentSessionStorage) GetExperimentWithoutCustomGroupsByNamespace(namespace string) dto.NameSpaceExperiments {
	r.mu.RLock()
	defer r.mu.RUnlock()

	experiments := r.sessionWithoutCustomParamGroups[namespace]

	return experiments
}

func (r *RawExperimentSessionStorage) GetFeatureTogglesByNamespace(namespace string) dto.NamespaceFeatureToggle {
	r.mu.RLock()
	defer r.mu.RUnlock()

	experiments := r.sessionFeatureToggle[namespace]

	return experiments
}
