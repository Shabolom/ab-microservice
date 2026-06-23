package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"ab/pkg/utils"
	"context"

	"go.uber.org/zap"
)

func (s *Service) UpdateRolloutPercentage(ctx context.Context, update *dto.FeatureTogglePercentageUpdate) error {
	s.logger.Info(
		"update feature toggle rollout percentage started",
		zap.Int64("id", update.FeatureID),
		zap.Any("rollout_percentage", update.Percentage),
		zap.Any("ios", update.Ios),
		zap.Any("android", update.Android),
		zap.Any("web", update.Web),
	)

	if update.FeatureID <= 0 {
		s.logger.Warn(
			"update feature toggle rollout percentage validation failed",
			zap.Int64("id", update.FeatureID),
			zap.Error(shortcut.ErrFeatureToggleIDRequired),
		)

		return shortcut.ErrFeatureToggleIDRequired
	}

	if err := utils.ValidatePercentage(update.Percentage); err != nil {
		return err
	}

	if err := utils.ValidatePercentage(update.Ios); err != nil {
		return err
	}

	if err := utils.ValidatePercentage(update.Android); err != nil {
		return err
	}

	if err := utils.ValidatePercentage(update.Web); err != nil {
		return err
	}

	feature, err := s.featureTogglesRepo.GetByID(ctx, update.FeatureID)
	if err != nil {
		s.logger.Error(
			"failed to get feature toggle",
			zap.Int64("id", update.FeatureID),
			zap.Error(err),
		)

		return err
	}

	err = s.updateRolloutPercentage(ctx, feature, update)
	if err != nil {
		s.logger.Error(
			"update feature toggle rollout percentage failed",
			zap.Int64("id", update.FeatureID),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"update feature toggle rollout percentage finished",
		zap.Int64("id", update.FeatureID),
	)

	return nil
}

func (s *Service) updateRolloutPercentage(
	ctx context.Context,
	feature *dto.RawFeatureToggle,
	update *dto.FeatureTogglePercentageUpdate,
) error {
	switch feature.Status {
	case dto.FeatureToggleStatusArchived,
		dto.FeatureToggleStatusDisabled,
		dto.FeatureToggleStatusDraft,
		dto.FeatureToggleStatusActive:
		err := s.featureTogglesRepo.UpdatePercentage(ctx, update)
		if err != nil {
			return err
		}
		return nil

	default:
		return shortcut.ErrFeatureToggleInvalidStatus
	}
	
}
