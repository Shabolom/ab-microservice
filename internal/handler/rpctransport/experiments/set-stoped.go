package experiments

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) SetStopedExperiment(ctx context.Context, req *authv1.SetStopedExperimentRequest) (*authv1.StockReply, error) {
	expID := req.GetExperimentId()

	err := h.experimentService.SetStopedStatus(ctx, expID)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("experiment was stoped, experiment id: %v", expID),
	}, nil
}
