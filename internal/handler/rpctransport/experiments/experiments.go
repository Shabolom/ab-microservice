package experiments

import (
	"ab/internal/dto"
)

type experimentService interface {
	GetExperiments(parameters *dto.RequestParameters) ([]*dto.GetExperimentsReply, error)
}
type Handler struct {
	experimentService experimentService
}

func New(experimentService experimentService) *Handler {
	return &Handler{
		experimentService: experimentService,
	}
}
