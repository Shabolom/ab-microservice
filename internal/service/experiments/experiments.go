package experiments

import (
	"ab/internal/dto"
	"ab/internal/dto/kafka-messege-dto"
	"context"

	"go.uber.org/zap"
)

type inMemoryStorage interface {
	GetExperimentByNamespace(namespace string) dto.NameSpaceExperiments
}

type nameSpaceRepo interface {
	GetList(ctx context.Context) ([]dto.NameSpace, error)
	GetByID(ctx context.Context, namespaceID int64) (*dto.NameSpace, error)
}

type groupRepo interface {
	GetListByExperimentID(ctx context.Context, experimentID string) ([]*dto.Group, error)
	Update(ctx context.Context, group *dto.UpdateGroup) (*dto.Group, error)
	GetByID(ctx context.Context, id int64) (*dto.Group, error)
	Post(ctx context.Context, experimentID int64, group *dto.Group) (*dto.Group, error)
	GetList(ctx context.Context, limit int, id int) ([]*dto.Group, error)
}

type experimentRepo interface {
	GetExperiment(ctx context.Context, id int64) (*dto.Experiment, error)
	CreateExperiment(ctx context.Context, experiment *dto.Experiment) (*dto.Experiment, error)
	GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error)
	GetLayerBucketsInPeriod(ctx context.Context, experiment *dto.Experiment) ([]dto.LayerBuckets, error)
	SetReadyStatus(ctx context.Context, experimentID int64, status string, layerBuckets []dto.LayerBuckets) error
	SetStatusStopped(ctx context.Context, expID int64) error
}

type layerRepo interface {
	GetLayers(ctx context.Context, layerIDs []int64) ([]dto.Layer, error)
}

type kafkaProducer interface {
	WriteEvent(event *kafkaMessageDto.UserInExperimentMessage) error
}
type Service struct {
	groupRepo       groupRepo
	experimentRepo  experimentRepo
	nameSpaceRepo   nameSpaceRepo
	inMemoryStorage inMemoryStorage
	layerRepo       layerRepo
	kafkaProducer   kafkaProducer
	logger          *zap.Logger
}

func New(
	groupRepo groupRepo,
	experimentRepo experimentRepo,
	nameSpaceRepo nameSpaceRepo,
	inMemoryStorage inMemoryStorage,
	layerRepo layerRepo,
	kafkaProducer kafkaProducer,
	logger *zap.Logger,
) *Service {
	return &Service{
		groupRepo:       groupRepo,
		experimentRepo:  experimentRepo,
		nameSpaceRepo:   nameSpaceRepo,
		inMemoryStorage: inMemoryStorage,
		layerRepo:       layerRepo,
		kafkaProducer:   kafkaProducer,
		logger:          logger,
	}
}
