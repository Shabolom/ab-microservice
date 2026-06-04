package inMemmoryCashe

import (
	"ab/internal/dto"
	"fmt"
)

func (r *RawExperimentSessionStorage) GetExperimentByNamespace(namespace string) dto.NameSpaceExperiments {
	r.mu.RLock()
	defer r.mu.RUnlock()
	fmt.Println("GetExperimentByNamespace 123123123", namespace)
	experiments := r.session[namespace]
	fmt.Println("GetExperimentByNamespace 44444444", experiments)
	return experiments
}
