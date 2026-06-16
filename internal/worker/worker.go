package worker

import (
	"ab/internal/dto"
	"context"
	"time"

	"go.uber.org/zap"
)

type metrics interface {
	SetActiveExperimentsWithGroup(count int)
	SetActiveExperimentsWithoutGroup(count int)
	ObserveCacheRefresh(startedAt time.Time)
}

type namespaceRepository interface {
	GetList(ctx context.Context) ([]dto.NameSpace, error)
}

type experimentRepository interface {
	GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error)
	UpdateReadyToStart(ctx context.Context, startDate time.Time) ([]int64, error)
	UpdateExpired(ctx context.Context, date time.Time) ([]int64, error)
}

type rawExperimentCache interface {
	ReplaceWithCustomGroups(newCash map[string]dto.NameSpaceExperiments)
	ReplaceWithoutCustomGroups(newCash map[string]dto.NameSpaceExperiments)
}
type Worker struct {
	namespaceRepository  namespaceRepository
	experimentRepository experimentRepository
	rawExperimentCache   rawExperimentCache
	metrics              metrics
	ctxInterval          int
	operatingInterval    int
	logger               *zap.Logger
}

func New(
	namespaceRepository namespaceRepository,
	experimentRepository experimentRepository,
	rawExperimentCache rawExperimentCache,
	metrics metrics,
	ctxInterval int,
	operatingInterval int,
	logger *zap.Logger,
) *Worker {
	return &Worker{
		namespaceRepository:  namespaceRepository,
		experimentRepository: experimentRepository,
		rawExperimentCache:   rawExperimentCache,
		metrics:              metrics,
		ctxInterval:          ctxInterval,
		operatingInterval:    operatingInterval,
		logger:               logger,
	}
}
