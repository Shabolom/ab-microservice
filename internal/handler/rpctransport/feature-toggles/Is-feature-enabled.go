package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) IsFeatureEnabled(ctx context.Context, req *authv1.IsFeatureEnabledRequest) (*authv1.StockReply, error) {
	id := req.GetFeatureToggleId()

	err := h.featureTogglesService.IsEnable(ctx, id)
	if err != nil {
		fmt.Println(err, 123123)
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("feature toggle is enable id %d", id),
	}, nil
}
