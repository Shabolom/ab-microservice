package experiments

import (
	"ab/internal/dto"
	"context"
)

type experimentService interface {
	GetExperiments(ctx context.Context, splitID int64) ([]*dto.GetExperimentsReply, error)
}
type Handler struct {
	experimentService experimentService
}

func New(experimentService experimentService) *Handler {
	return &Handler{
		experimentService: experimentService,
	}
}
