package customParams

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) GetCustomParamByID(ctx context.Context, req *authv1.GetCustomParamByIDRequest) (*authv1.GetCustomParamByIDReply, error) {
	id := req.GetId()

	customParam, err := h.customParamsService.GetById(ctx, id)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.GetCustomParamByIDReply{
		ErrInfoReason: authv1.GetCustomParamByIDReply_STATUS_OK,
		Message:       fmt.Sprintf("custom param id %v :", id),
		CustomParam: &authv1.CustomParam{
			Id:          customParam.ID,
			NamespaceId: customParam.NameSpaceID,
			Name:        customParam.Name,
			Type:        customParam.Type,
		},
	}, nil
}
