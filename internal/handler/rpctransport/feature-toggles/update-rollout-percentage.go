package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) UpdateFeatureToggleRollout(ctx context.Context, req *authv1.UpdateFeatureToggleRolloutRequest) (*authv1.StockReply, error) {
	id := int(req.GetFeatureToggleId())
	rolloutPercentage := int(req.GetRolloutPercentage())

	err := h.featureTogglesService.UpdateRolloutPercentage(ctx, id, rolloutPercentage)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("Feature toggle rollout set to %d", id),
	}, nil
}
