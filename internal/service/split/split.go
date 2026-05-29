package split

import (
	"ab/internal/dto"
	"context"
)

type groupRepo interface {
	GetListByExperimentID(ctx context.Context, experimentID string) ([]*dto.Group, error)
	Update(ctx context.Context, group *dto.UpdateGroup) (*dto.Group, error)
	GetByID(ctx context.Context, id int64) (*dto.Group, error)
	Post(ctx context.Context, experimentID string, group *dto.Group) (*dto.Group, error)
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context, limit int, id int) ([]*dto.Group, error)
}

type experimentRepo interface {
	GetExperiment(ctx context.Context, id string) (*dto.Experiment, error)
	CreateExperiment(ctx context.Context, experiment *dto.Experiment) (*dto.Experiment, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, req *dto.UpdateExperiment) (*dto.Experiment, error)
	GetList(ctx context.Context) ([]*dto.Experiment, error)
}
type Service struct {
	groupRepo      groupRepo
	experimentRepo experimentRepo
}

func New(groupRepo groupRepo, experimentRepo experimentRepo) *Service {
	return &Service{
		groupRepo:      groupRepo,
		experimentRepo: experimentRepo,
	}
}
