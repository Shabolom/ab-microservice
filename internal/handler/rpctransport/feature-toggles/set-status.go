package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) SetFeatureToggleStatus(ctx context.Context, req *authv1.SetFeatureToggleStatusRequest) (*authv1.StockReply, error) {
	id := req.GetFeatureToggleId()
	status := req.GetStatus()

	err := h.featureTogglesService.SetStatus(ctx, id, status)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("set status %v for feature toggle %d", status, id),
	}, nil
}
