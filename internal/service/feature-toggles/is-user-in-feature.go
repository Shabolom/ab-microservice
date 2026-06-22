package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/utils"
	"context"
)

func (s *Service) IsUserInFeature(ctx context.Context, splitID int64, namespace string) ([]*dto.FeatureReply, error) {
	featureByNamespace := s.inMemoryStorage.GetFeatureTogglesByNamespace(namespace)
	answer := make([]*dto.FeatureReply, 0)

	for _, feature := range featureByNamespace.FeatureToggles {

		if feature.Status != dto.FeatureToggleStatusActive {
			continue
		}

		if !utils.IsUserInRolloutPercentage(feature.ID, splitID, feature.Buckets) {
			continue
		}

		answer = append(answer, &dto.FeatureReply{
			FeatureID:   feature.ID,
			FeatureName: feature.Name,
		})
	}

	return answer, nil
}
