package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) UpdateFeatureToggleRollout(ctx context.Context, req *authv1.UpdateFeatureToggleRolloutRequest) (*authv1.StockReply, error) {
	id := req.GetFeatureToggleId()

	err := h.featureTogglesService.UpdateRolloutPercentage(ctx, &dto.FeatureTogglePercentageUpdate{
		FeatureID:  id,
		Percentage: req.RolloutPercentage,
		Ios:        req.IosRolloutPercentage,
		Android:    req.AndroidRolloutPercentage,
		Web:        req.WebRolloutPercentage,
	})
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("feature toggle rollout set to %d", id),
	}, nil
}
