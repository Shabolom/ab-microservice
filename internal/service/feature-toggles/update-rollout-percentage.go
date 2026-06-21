package featureToggles

import (
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) UpdateRolloutPercentage(ctx context.Context, id int, percentage int) error {
	s.logger.Info(
		"update feature toggle rollout percentage started",
		zap.Int("id", id),
		zap.Int("rollout_percentage", percentage),
	)

	switch {
	case id <= 0:
		s.logger.Warn(
			"update feature toggle rollout percentage validation failed",
			zap.Int("id", id),
			zap.Int("rollout_percentage", percentage),
			zap.Error(shortcut.ErrValidation),
		)

		return shortcut.ErrValidation

	case percentage < 0 || percentage > 100:
		s.logger.Warn(
			"update feature toggle rollout percentage validation failed",
			zap.Int("id", id),
			zap.Int("rollout_percentage", percentage),
			zap.Error(shortcut.ErrFeatureToggleRolloutOutOfRange),
		)

		return shortcut.ErrFeatureToggleRolloutOutOfRange
	}

	err := s.featureTogglesRepo.UpdateRolloutPercentage(
		ctx,
		id,
		percentage,
	)
	if err != nil {
		s.logger.Error(
			"update feature toggle rollout percentage failed",
			zap.Int("id", id),
			zap.Int("rollout_percentage", percentage),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"update feature toggle rollout percentage finished",
		zap.Int("id", id),
		zap.Int("rollout_percentage", percentage),
	)

	return nil
}
