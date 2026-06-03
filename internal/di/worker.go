package di

import "ab/internal/worker"

func (d *DI) GetWorker() *worker.Worker {
	return worker.New(
		d.GetNamespacePgRepo(),
		d.GetExperimentPgRepo(),
		d.GetInMemoryCache(),
		d.Logger(),
	)
}
