package worker

import "context"

func (w *Worker) Start(ctx context.Context) {
	go w.startExperimentsCacheWorker(ctx)
	go w.startActivateExperimentWorker(ctx)
	go w.startExpiredExperimentWorker(ctx)
	go w.startFeatureToggleCacheWorker(ctx)
}
