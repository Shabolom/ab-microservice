package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) CreateFeatureToggle(ctx context.Context, req *authv1.CreateFeatureToggleRequest) (*authv1.StockReply, error) {
	featureToggle := &dto.FeatureToggle{
		NamespaceID:       req.GetNamespaceId(),
		Name:              req.GetName(),
		RolloutPercentage: req.RolloutPercentage,
		IOS:               req.IosRolloutPercentage,
		Android:           req.AndroidRolloutPercentage,
		Web:               req.WebRolloutPercentage,
	}

	err := h.featureTogglesService.Create(ctx, featureToggle)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("feature toggle %v with id %d was created", featureToggle.Name, featureToggle.ID),
	}, nil
}
