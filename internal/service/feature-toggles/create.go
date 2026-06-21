package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, featureToggle *dto.FeatureToggle) error {
	s.logger.Info(
		"create feature toggle started",
		zap.String("name", featureToggle.Name),
		zap.Int64("namespace_id", featureToggle.NamespaceID),
	)

	featureToggle.Status = dto.FeatureToggleStatusDraft

	err := s.createValidation(featureToggle)
	if err != nil {
		s.logger.Warn(
			"create feature toggle validation failed",
			zap.String("name", featureToggle.Name),
			zap.Int64("namespace_id", featureToggle.NamespaceID),
			zap.Int("rollout_percentage", featureToggle.RolloutPercentage),
			zap.Error(err),
		)

		return err
	}

	err = s.featureTogglesRepo.Create(ctx, featureToggle)
	if err != nil {
		s.logger.Error(
			"create feature toggle failed",
			zap.String("name", featureToggle.Name),
			zap.Int64("namespace_id", featureToggle.NamespaceID),
			zap.Int("rollout_percentage", featureToggle.RolloutPercentage),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"create feature toggle finished",
		zap.Int64("id", featureToggle.ID),
		zap.String("name", featureToggle.Name),
		zap.Int64("namespace_id", featureToggle.NamespaceID),
	)

	return nil
}

func (s *Service) createValidation(featureToggle *dto.FeatureToggle) error {
	switch {
	case featureToggle.Name == "":
		return shortcut.ErrFeatureToggleNameRequired
	case featureToggle.NamespaceID == 0 || featureToggle.NamespaceID < 0:
		return shortcut.ErrFeatureToggleNamespaceIDRequired
	case featureToggle.RolloutPercentage < 0 || featureToggle.RolloutPercentage > 100:
		return shortcut.ErrFeatureToggleRolloutOutOfRange
	}

	return nil
}
