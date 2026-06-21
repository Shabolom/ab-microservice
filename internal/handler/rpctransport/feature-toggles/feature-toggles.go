package featureToggles

import (
	"ab/internal/dto"
	"context"
)

type featureTogglesService interface {
	Create(ctx context.Context, featureToggle *dto.FeatureToggle) error
	SetStatus(ctx context.Context, id int, status string) error
	UpdateRolloutPercentage(ctx context.Context, id int, percentage int) error
}
type Handler struct {
	featureTogglesService featureTogglesService
}

func New(featureTogglesService featureTogglesService) *Handler {
	return &Handler{
		featureTogglesService: featureTogglesService,
	}
}
