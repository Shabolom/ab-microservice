package worker

import "context"

func (w *Worker) Start(ctx context.Context) {
	go func() {
		w.startExperimentsCacheWorker(ctx)
	}()
	go func() {
		w.logger.Info("activate experiment worker started")
		w.startActivateExperimentWorker(ctx)
	}()
}
