package experiments

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) SetReadyExperiment(ctx context.Context, req *authv1.SetReadyExperimentRequest) (*authv1.StockReply, error) {
	err := h.experimentService.SetReady(ctx, req.GetExperimentId())
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("Set ready experiment success status experiment id: %v", req.GetExperimentId()),
	}, nil
}
