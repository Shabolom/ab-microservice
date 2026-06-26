package featureToggles

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type featureTogglesRepo interface {
	Create(ctx context.Context, featureToggle *dto.FeatureToggle) error
	IsFeatureTogglesEnable(ctx context.Context, id int64) (string, error)
	GetActiveByNamespaceID(ctx context.Context, id int64) ([]dto.RawFeatureToggle, error)
	GetByID(ctx context.Context, id int64) (*dto.RawFeatureToggle, error)
	SetStatus(ctx context.Context, id int64, status string) error
	UpdatePercentage(ctx context.Context, update *dto.FeatureTogglePercentageUpdate) error
}

type inMemoryStorage interface {
	GetFeatureTogglesByNamespace(namespace string) dto.NamespaceFeatureToggle
}

type Service struct {
	featureTogglesRepo featureTogglesRepo
	inMemoryStorage    inMemoryStorage
	logger             *zap.Logger
}

func New(
	featureTogglesRepo featureTogglesRepo,
	inMemoryStorage inMemoryStorage,
	logger *zap.Logger,
) *Service {
	return &Service{
		logger:             logger,
		inMemoryStorage:    inMemoryStorage,
		featureTogglesRepo: featureTogglesRepo,
	}
}
