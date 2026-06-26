package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) IsEnable(ctx context.Context, id int64) error {
	s.logger.Info(
		"check feature toggle status started",
		zap.Int64("id", id),
	)

	status, err := s.featureTogglesRepo.IsFeatureTogglesEnable(ctx, id)
	if err != nil {
		s.logger.Warn(
			"failed to get feature toggle status",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return err
	}

	if status != dto.FeatureToggleStatusActive {
		s.logger.Warn(
			"feature toggle is not active",
			zap.Int64("id", id),
			zap.String("status", status),
		)

		return shortcut.ErrFeatureToggleNotActive
	}

	s.logger.Info(
		"feature toggle is active",
		zap.Int64("id", id),
	)

	return nil
}
