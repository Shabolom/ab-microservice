package cache

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) StartExperimentsCacheWorker(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	w.logger.Info("cache cache started")

	if err := w.refresh(ctx); err != nil {
		w.logger.Warn("failed to refresh experiment cache", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("cache cache stopped")
			return

		case <-ticker.C:
			if err := w.refresh(ctx); err != nil {
				w.logger.Error("failed to refresh experiment cache", zap.Error(err))
			}
		}
	}
}
