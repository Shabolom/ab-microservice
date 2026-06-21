package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Service) IsEnable(ctx context.Context, id int64) error {
	statuses, err := s.featureTogglesRepo.IsFeatureTogglesEnable(ctx, id)
	if err != nil {
		return err
	}

	if statuses != dto.FeatureToggleStatusActive {
		return shortcut.ErrFeatureToggleNotActive
	}

	return nil
}
