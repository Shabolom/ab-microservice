package worker

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type nameSpaceRepository interface {
	GetList(ctx context.Context) ([]dto.NameSpace, error)
}

type experimentRepository interface {
	GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error)
	GetLayersWithExperiments(ctx context.Context) ([]dto.LayerWithExperiments, error)
	UpdateStatusBuckets(ctx context.Context, experimentID int64, status string, buckets []int64) error
}

type rawExperimentCache interface {
	Replace(newCash map[string]dto.NameSpaceExperiments)
}
type Worker struct {
	nameSpaceRepository  nameSpaceRepository
	experimentRepository experimentRepository
	rawExperimentCache   rawExperimentCache
	logger               *zap.Logger
}

func New(
	nameSpaceRepository nameSpaceRepository,
	experimentRepository experimentRepository,
	rawExperimentCache rawExperimentCache,
	logger *zap.Logger,
) *Worker {
	return &Worker{
		nameSpaceRepository:  nameSpaceRepository,
		experimentRepository: experimentRepository,
		rawExperimentCache:   rawExperimentCache,
		logger:               logger,
	}
}
