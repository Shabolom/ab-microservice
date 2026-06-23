package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
)

func (h *Handler) IsUserInFeature(ctx context.Context, req *authv1.IsUserInFeatureRequest) (*authv1.IsUserInFeatureReply, error) {
	reqInfo := &dto.UserInFeatureReq{
		UserId:    req.GetUserId(),
		Namespace: req.GetNamespace(),
		Platform:  req.GetPlatform(),
	}

	features, err := h.featureTogglesService.IsUserInFeature(ctx, reqInfo)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var answer []*authv1.Feature

	for _, feature := range features {
		answer = append(answer, &authv1.Feature{
			FeatureId:   feature.FeatureID,
			FeatureName: feature.FeatureName,
		})
	}

	return &authv1.IsUserInFeatureReply{
		ErrInfoReason: authv1.IsUserInFeatureReply_STATUS_OK,
		Features:      answer,
	}, nil
}
