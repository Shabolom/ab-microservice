package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetFeatureToggleByID(ctx context.Context, req *authv1.GetFeatureToggleByIDRequest) (*authv1.GetFeatureToggleByIDReply, error) {
	id := req.GetId()

	feature, err := h.featureTogglesService.GetByID(ctx, id)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	featureToggleResponse := &authv1.FeatureToggle{
		Id:          feature.ID,
		NamespaceId: feature.NamespaceID,
		Name:        feature.Name,
		Status:      feature.Status,
		CreatedAt:   timestamppb.New(feature.CreatedAt),
		UpdatedAt:   timestamppb.New(feature.UpdatedAt),
	}

	if feature.RolloutPercentage != nil {
		featureToggleResponse.RolloutPercentage = feature.RolloutPercentage
	}

	if feature.Web != nil {
		featureToggleResponse.Web = feature.Web
	}

	if feature.IOS != nil {
		featureToggleResponse.Ios = feature.IOS
	}

	if feature.Android != nil {
		featureToggleResponse.Android = feature.Android
	}

	if feature.DeletedAt != nil {
		featureToggleResponse.DeletedAt = timestamppb.New(*feature.DeletedAt)
	}

	return &authv1.GetFeatureToggleByIDReply{
		ErrInfoReason: authv1.GetFeatureToggleByIDReply_STATUS_OK,
		Message:       fmt.Sprintf("featureToggle with id %v :", id),
		FeatureToggle: featureToggleResponse,
	}, nil
}
