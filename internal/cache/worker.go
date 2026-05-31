package cache

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
