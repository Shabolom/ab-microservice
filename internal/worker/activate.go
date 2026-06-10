package worker

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) startActivateExperimentWorker(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(w.operatingInterval) * time.Second)
	defer ticker.Stop()

	if err := w.activateExperimentWorker(ctx); err != nil {
		w.logger.Error(
			"activate experiment iteration failed",
			zap.Error(err),
		)
	}

	for {
		select {
		case <-ctx.Done():
			w.logger.Info(
				"activate experiment worker stopped",
				zap.Error(ctx.Err()),
			)
			return

		case <-ticker.C:
			if err := w.activateExperimentWorker(ctx); err != nil {
				w.logger.Error(
					"activate experiment iteration failed",
					zap.Error(err),
				)
			}
		}
	}
}

func (w *Worker) activateExperimentWorker(ctx context.Context) error {
	now := time.Now()

	activatedExp, err := w.experimentRepository.UpdateReadyToStart(ctx, now)
	if err != nil {
		return err
	}

	if len(activatedExp) > 0 {
		w.logger.Info(
			"experiments activated",
			zap.Int("count", len(activatedExp)),
		)
	}

	return nil
}
