package featureToggles

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetFeatureToggles(ctx context.Context, req *emptypb.Empty) (*authv1.GetFeatureToggleBysReply, error) {
	features, err := h.featureTogglesService.GetList(ctx)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var featuresResponse []*authv1.FeatureToggle
	for _, feature := range features {
		featureResponse := &authv1.FeatureToggle{
			Id:          feature.ID,
			NamespaceId: feature.NamespaceID,
			Name:        feature.Name,
			Status:      feature.Status,
			CreatedAt:   timestamppb.New(feature.CreatedAt),
			UpdatedAt:   timestamppb.New(feature.UpdatedAt),
		}

		if feature.RolloutPercentage != nil {
			featureResponse.RolloutPercentage = feature.RolloutPercentage
		}

		if feature.Web != nil {
			featureResponse.Web = feature.Web
		}

		if feature.IOS != nil {
			featureResponse.Ios = feature.IOS
		}

		if feature.Android != nil {
			featureResponse.Android = feature.Android
		}

		if feature.DeletedAt != nil {
			featureResponse.DeletedAt = timestamppb.New(*feature.DeletedAt)
		}

		featuresResponse = append(featuresResponse, featureResponse)
	}

	return &authv1.GetFeatureToggleBysReply{
		ErrInfoReason:  authv1.GetFeatureToggleBysReply_STATUS_OK,
		Message:        "feature toggles:",
		FeatureToggles: featuresResponse,
	}, nil
}
