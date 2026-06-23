package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/utils"
	"context"
)

func (s *Service) IsUserInFeature(ctx context.Context, req *dto.UserInFeatureReq) ([]*dto.FeatureReply, error) {
	featureByNamespace := s.inMemoryStorage.GetFeatureTogglesByNamespace(req.Namespace)
	answer := make([]*dto.FeatureReply, 0)

	for _, feature := range featureByNamespace.FeatureToggles {
		if feature.Status != dto.FeatureToggleStatusActive {
			continue
		}

		rolloutPercentage := getRolloutPercentageByPlatform(feature, req.Platform)
		if rolloutPercentage == nil {
			continue
		}

		if !utils.IsUserInRolloutPercentageWithoutBuckets(feature.ID, req.UserId, *rolloutPercentage) {
			continue
		}

		answer = append(answer, &dto.FeatureReply{
			FeatureID:   feature.ID,
			FeatureName: feature.Name,
		})
	}

	return answer, nil
}

func getRolloutPercentageByPlatform(feature dto.RawFeatureToggle, platform string) *int64 {
	switch platform {
	case dto.FeatureToggleIosPlatform:
		if feature.IOS != nil {
			return feature.IOS
		}
		return nil
	case dto.FeatureToggleAndroidPlatform:
		if feature.Android != nil {
			return feature.Android
		}
		return nil
	case dto.FeatureToggleWebPlatform:
		if feature.Web != nil {
			return feature.Web
		}
		return nil
	}

	return feature.RolloutPercentage
}
