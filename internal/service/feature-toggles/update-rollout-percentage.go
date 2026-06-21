package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"ab/pkg/utils"
	"context"

	"go.uber.org/zap"
)

func (s *Service) UpdateRolloutPercentage(ctx context.Context, id int64, percentage int64) error {
	s.logger.Info(
		"update feature toggle rollout percentage started",
		zap.Int64("id", id),
		zap.Int64("rollout_percentage", percentage),
	)

	switch {
	case id <= 0:
		s.logger.Warn(
			"update feature toggle rollout percentage validation failed",
			zap.Int64("id", id),
			zap.Int64("rollout_percentage", percentage),
			zap.Error(shortcut.ErrFeatureToggleIDRequired),
		)

		return shortcut.ErrFeatureToggleIDRequired

	case percentage < 0 || percentage > 100:
		s.logger.Warn(
			"update feature toggle rollout percentage validation failed",
			zap.Int64("id", id),
			zap.Int64("rollout_percentage", percentage),
			zap.Error(shortcut.ErrFeatureToggleRolloutOutOfRange),
		)

		return shortcut.ErrFeatureToggleRolloutOutOfRange
	}

	feature, err := s.featureTogglesRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error(
			"failed to get feature toggle",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return err
	}

	if feature.RolloutPercentage == percentage {
		return shortcut.ErrFeatureToggleSameRolloutPercentage
	}

	err = s.updateRolloutPercentage(ctx, feature, percentage)
	if err != nil {
		s.logger.Error(
			"update feature toggle rollout percentage failed",
			zap.Int64("id", id),
			zap.Int64("rollout_percentage", percentage),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"update feature toggle rollout percentage finished",
		zap.Int64("id", id),
		zap.Int64("rollout_percentage", percentage),
	)

	return nil
}

func (s *Service) updateRolloutPercentage(ctx context.Context, feature *dto.RawFeatureToggle, newRolloutPercentage int64) error {
	switch feature.Status {
	case dto.FeatureToggleStatusActive:
		activeFeaturesByNamespace, err := s.featureTogglesRepo.GetActiveByNamespaceID(ctx, feature.NamespaceID)
		if err != nil {
			return err
		}

		usedBuckets := make([]int64, 0, 100)
		for _, activeFeature := range activeFeaturesByNamespace {
			usedBuckets = append(usedBuckets, activeFeature.Buckets...)
		}

		if newRolloutPercentage > feature.RolloutPercentage {
			difference := newRolloutPercentage - feature.RolloutPercentage

			if int64(len(usedBuckets))+difference > 100 {
				return shortcut.ErrFeatureToggleNotEnoughBuckets
			}

			newBuckets := utils.GenerateBuckets(usedBuckets, difference)
			feature.Buckets = append(feature.Buckets, newBuckets...)
		} else {
			feature.Buckets = feature.Buckets[:int(newRolloutPercentage)]
		}

		err = s.featureTogglesRepo.UpdatePercentageAndBuckets(ctx, feature.ID, feature.Buckets)
		if err != nil {
			return err
		}

		return nil

	case
		dto.FeatureToggleStatusArchived,
		dto.FeatureToggleStatusDisabled,
		dto.FeatureToggleStatusDraft:
		err := s.featureTogglesRepo.UpdatePercentage(ctx, feature.ID, newRolloutPercentage)
		if err != nil {
			return err
		}

		return nil

	default:
		return shortcut.ErrFeatureToggleInvalidStatus
	}
}
