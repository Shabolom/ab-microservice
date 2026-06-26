package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetByID(ctx context.Context, id int64) (*dto.RawFeatureToggle, error) {
	s.logger.Info("GetList featureToggle started")

	if id == 0 {
		s.logger.Warn("id is 0")
		return nil, shortcut.ErrValidation
	}

	feature, err := s.featureTogglesRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Warn("GetList featureToggle failed",
			zap.Error(err),
			zap.Int64("id", id),
		)

		return nil, err
	}

	s.logger.Info("GetList featureToggle started")
	return feature, nil
}
