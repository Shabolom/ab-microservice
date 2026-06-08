package experiments

import (
	"ab/internal/dto"
	"ab/internal/dto/kafka-messege-dto"
	"context"

	"go.uber.org/zap"
)

type inMemoryStorage interface {
	GetExperimentWithCustomGroupsByNamespace(namespace string) dto.NameSpaceExperiments
	GetExperimentWithoutCustomGroupsByNamespace(namespace string) dto.NameSpaceExperiments
}

type nameSpaceRepo interface {
	GetList(ctx context.Context) ([]dto.NameSpace, error)
	GetByID(ctx context.Context, namespaceID int64) (*dto.NameSpace, error)
	GetByName(ctx context.Context, name string) (*dto.NameSpace, error)
}

type groupRepo interface {
	GetListByExperimentID(ctx context.Context, experimentID string) ([]*dto.Group, error)
	Update(ctx context.Context, group *dto.UpdateGroup) (*dto.Group, error)
	GetByID(ctx context.Context, id int64) (*dto.Group, error)
	Post(ctx context.Context, experimentID int64, group *dto.Group) (*dto.Group, error)
	GetList(ctx context.Context, limit int, id int) ([]*dto.Group, error)
}

type experimentRepo interface {
	GetExperimentWithLayers(ctx context.Context, id int64) (*dto.Experiment, error)
	CreateExperiment(ctx context.Context, experiment *dto.Experiment) (*dto.Experiment, error)
	GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error)
	GetLayerBucketsInPeriod(ctx context.Context, experiment *dto.Experiment) ([]dto.LayerBuckets, error)
	SetReadyStatus(ctx context.Context, experimentID int64, status string, layerBuckets []dto.LayerBuckets) error
	SetStatusStopped(ctx context.Context, expID int64) error
	CreateAndGetLayerExperiment(ctx context.Context, layerIDs []int64, expID int64) ([]dto.LayerBuckets, error)
}

type layerRepo interface {
	GetLayers(ctx context.Context, layerIDs []int64) ([]dto.Layer, error)
	GetIDsByNamespaceId(ctx context.Context, namespaceId int64) ([]int64, error)
}

type kafkaProducer interface {
	WriteEvent(event *kafkaMessageDto.UserInExperimentMessage) error
}

type customParamRepo interface {
	GetById(ctx context.Context, id int64) (dto.CustomParams, error)
	CountDistinctNamespaces(ctx context.Context, ids []int64) (int64, error)
}

type Service struct {
	groupRepo       groupRepo
	experimentRepo  experimentRepo
	nameSpaceRepo   nameSpaceRepo
	inMemoryStorage inMemoryStorage
	layerRepo       layerRepo
	kafkaProducer   kafkaProducer
	customParamRepo customParamRepo
	logger          *zap.Logger
}

func New(
	groupRepo groupRepo,
	experimentRepo experimentRepo,
	nameSpaceRepo nameSpaceRepo,
	inMemoryStorage inMemoryStorage,
	layerRepo layerRepo,
	kafkaProducer kafkaProducer,
	customParamRepo customParamRepo,
	logger *zap.Logger,
) *Service {
	return &Service{
		groupRepo:       groupRepo,
		experimentRepo:  experimentRepo,
		nameSpaceRepo:   nameSpaceRepo,
		inMemoryStorage: inMemoryStorage,
		layerRepo:       layerRepo,
		kafkaProducer:   kafkaProducer,
		customParamRepo: customParamRepo,
		logger:          logger,
	}
}
