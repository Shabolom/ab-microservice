package experiments

import (
	"ab/internal/dto"
	"context"
)

type experimentService interface {
	GetExperiments(parameters *dto.RequestParameters) ([]*dto.GetExperimentsReply, error)
	Create(ctx context.Context, reqExp *dto.Experiment) (*dto.Experiment, error)
	SetReady(ctx context.Context, targetExpID int64) error
	SetStopedStatus(ctx context.Context, expID int64) error
	GetById(ctx context.Context, id int64) (*dto.Experiment, error)
	GetList(ctx context.Context) ([]*dto.Experiment, error)
}
type Handler struct {
	experimentService experimentService
}

func New(experimentService experimentService) *Handler {
	return &Handler{
		experimentService: experimentService,
	}
}
