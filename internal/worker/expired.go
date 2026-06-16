package worker

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) startExpiredExperimentWorker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	if err := w.expiredExperimentWorker(ctx); err != nil {
		w.logger.Error(
			"expired experiment iteration failed",
			zap.Error(err),
		)
	}

	for {
		select {
		case <-ctx.Done():
			w.logger.Info(
				"expired experiment worker stopped",
				zap.Error(ctx.Err()),
			)
			return

		case <-ticker.C:
			if err := w.expiredExperimentWorker(ctx); err != nil {
				w.logger.Error(
					"expired experiment iteration failed",
					zap.Error(err),
				)
			}
		}
	}
}

func (w *Worker) expiredExperimentWorker(ctx context.Context) error {
	now := time.Now()

	expiredExp, err := w.experimentRepository.UpdateExpired(ctx, now)
	if err != nil {
		return err
	}

	if len(expiredExp) > 0 {
		w.logger.Info(
			"experiments marked as ended and deleted from layers",
			zap.Int("count", len(expiredExp)),
		)
	}

	return nil
}
