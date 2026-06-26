package featureToggles

import (
	"ab/internal/dto"
	"context"
)

type featureTogglesService interface {
	Create(ctx context.Context, featureToggle *dto.FeatureToggle) error
	SetStatus(ctx context.Context, id int64, status string) error
	UpdateRolloutPercentage(ctx context.Context, update *dto.FeatureTogglePercentageUpdate) error
	IsEnable(ctx context.Context, id int64) error
	IsUserInFeature(ctx context.Context, req *dto.UserInFeatureReq) ([]*dto.FeatureReply, error)
}
type Handler struct {
	featureTogglesService featureTogglesService
}

func New(featureTogglesService featureTogglesService) *Handler {
	return &Handler{
		featureTogglesService: featureTogglesService,
	}
}
