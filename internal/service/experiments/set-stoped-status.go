package experiments

import (
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) SetStopedStatus(ctx context.Context, expID int64) error {
	if expID == 0 {
		s.logger.Warn(
			"empty experiment id",
			zap.Int64("experiment_id", expID),
		)

		return shortcut.ErrValidation
	}

	s.logger.Info(
		"stopping experiment",
		zap.Int64("experiment_id", expID),
	)

	err := s.experimentRepo.SetStatusStopped(ctx, expID)
	if err != nil {
		s.logger.Warn(
			"failed to stop experiment",
			zap.Int64("experiment_id", expID),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"experiment stopped",
		zap.Int64("experiment_id", expID),
	)

	return nil
}
