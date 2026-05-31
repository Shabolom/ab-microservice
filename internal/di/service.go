package di

import "ab/internal/service/experiments"

func (d *DI) GetExampleService() *experiments.Service {
	return experiments.New(
		d.GetGroupPgRepo(),
		d.GetExperimentPgRepo(),
		d.GetNamespacePgRepo(),
		d.GetInMemoryCache(),
		d.Logger(),
	)
}
