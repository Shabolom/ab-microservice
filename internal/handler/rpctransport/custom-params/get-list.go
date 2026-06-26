package customParams

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetCustomParams(ctx context.Context, req *emptypb.Empty) (*authv1.GetCustomParamBysReply, error) {
	customParams, err := h.customParamsService.GetList(ctx)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var customParamsResponse []*authv1.CustomParam
	for _, customParam := range customParams {
		customParamsResponse = append(customParamsResponse, &authv1.CustomParam{
			Id:          customParam.ID,
			NamespaceId: customParam.NameSpaceID,
			Name:        customParam.Name,
			Type:        customParam.Type,
		})
	}

	return &authv1.GetCustomParamBysReply{
		ErrInfoReason: authv1.GetCustomParamBysReply_STATUS_OK,
		Message:       "custom params:",
		CustomParams:  customParamsResponse,
	}, nil
}
