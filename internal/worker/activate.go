package worker

import (
	"ab/pkg/shortcut"
	"context"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) startActivateExperimentWorker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
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

	readyToStartExp, err := w.experimentRepository.GetReadyToStart(ctx, now)
	if err != nil {
		return err
	}

	if len(readyToStartExp) > 0 {
		w.logger.Info(
			"experiments ready to activate found",
			zap.Int("count", len(readyToStartExp)),
		)
	}

	for _, expID := range readyToStartExp {
		err = w.experimentRepository.UpdateStatus(
			ctx,
			expID,
			shortcut.ExpStatusActive,
		)
		if err != nil {
			w.logger.Error(
				"failed to mark experiment as active",
				zap.Int64("experiment_id", expID),
				zap.Error(err),
			)
			continue
		}

		w.logger.Info(
			"experiment marked as active",
			zap.Int64("experiment_id", expID),
		)
	}

	return nil
}
