package worker

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) startExperimentsCacheWorker(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	w.logger.Info("worker worker started")

	if err := w.refresh(ctx); err != nil {
		w.logger.Warn("failed to refresh experiment worker", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("worker worker stopped")
			return

		case <-ticker.C:
			if err := w.refresh(ctx); err != nil {
				w.logger.Error("failed to refresh experiment worker", zap.Error(err))
			}
		}
	}
}
