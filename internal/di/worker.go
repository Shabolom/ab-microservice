package di

import "ab/internal/cache"

func (d *DI) GetWorker() *cache.Worker {
	return cache.New(
		d.GetNamespacePgRepo(),
		d.GetExperimentPgRepo(),
		d.GetInMemoryCache(),
		d.Logger(),
	)
}
