package featureToggles

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type featureTogglesRepo interface {
	Create(ctx context.Context, featureToggle *dto.FeatureToggle) error
	UpdatePercentageAndBuckets(ctx context.Context, id int64, buckets []int64) error
	IsFeatureTogglesEnable(ctx context.Context, id int64) (string, error)
	GetActiveByNamespaceID(ctx context.Context, id int64) ([]dto.RawFeatureToggle, error)
	GetByID(ctx context.Context, id int64) (*dto.RawFeatureToggle, error)
	SetStatusActive(ctx context.Context, id int64, buckets []int64) error
	SetStatusDisabled(ctx context.Context, id int64) error
	SetStatusArchived(ctx context.Context, id int64) error
	SetStatusDraft(ctx context.Context, id int64) error
	UpdatePercentage(ctx context.Context, id int64, percentage int64) error
}

type inMemoryStorage interface {
	GetFeatureTogglesByNamespace(namespace string) dto.NamespaceFeatureToggle
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
