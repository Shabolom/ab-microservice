package featureToggles

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetList(ctx context.Context) ([]*dto.RawFeatureToggle, error) {
	s.logger.Info("GetList featureToggle started")

	features, err := s.featureTogglesRepo.GetList(ctx)
	if err != nil {
		s.logger.Warn("GetList error",
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList featureToggle finished")
	return features, nil
}
