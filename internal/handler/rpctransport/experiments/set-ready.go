package experiments

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"context"
	"time"
)

func (h *Handler) SetReadyExperiment(context.Context, *authv1.SetReadyExperimentRequest) (*authv1.StockReply, error) {
	experimentInfo := &dto.ExperimentStatus{
		ExpID:     0,
		LayerID:   0,
		StartedAt: time.Time{},
		EndedAt:   time.Time{},
	}

	return &authv1.StockReply{}, nil
}
