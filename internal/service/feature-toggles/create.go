package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/utils"
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

	err := utils.CreateValidation(featureToggle)
	if err != nil {
		s.logger.Warn(
			"create feature toggle validation failed",
			zap.String("name", featureToggle.Name),
			zap.Int64("namespace_id", featureToggle.NamespaceID),
			zap.Any("rollout_percentage", featureToggle.RolloutPercentage),
			zap.Any("ios", featureToggle.IOS),
			zap.Any("android", featureToggle.Android),
			zap.Any("web", featureToggle.Web),
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
			zap.Any("rollout_percentage", featureToggle.RolloutPercentage),
			zap.Any("ios", featureToggle.IOS),
			zap.Any("android", featureToggle.Android),
			zap.Any("web", featureToggle.Web),
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
