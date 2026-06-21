package featureToggles

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type featureTogglesRepo interface {
	Create(ctx context.Context, featureToggle *dto.FeatureToggle) error
	SetStatus(ctx context.Context, id int, status string) error
	UpdateRolloutPercentage(ctx context.Context, id int, rolloutPercentage int) error
}

type Service struct {
	featureTogglesRepo featureTogglesRepo
	logger             *zap.Logger
}

func New(
	featureTogglesRepo featureTogglesRepo,
	logger *zap.Logger,
) *Service {
	return &Service{
		logger:             logger,
		featureTogglesRepo: featureTogglesRepo,
	}
}
