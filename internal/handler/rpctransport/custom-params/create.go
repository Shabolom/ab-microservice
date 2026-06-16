package customParams

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
	"fmt"
	"strings"
)

func (h *Handler) CreateCustomParam(ctx context.Context, req *authv1.CreateCustomParamRequest) (*authv1.StockReply, error) {
	param := &dto.CustomParams{
		Name:        req.GetName(),
		NameSpaceID: req.GetNamespaceId(),
		Type:        strings.ToUpper(req.GetType()),
	}

	param, err := h.customParamsService.Create(ctx, param)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("custom param %v successfully, paramID: %v", param.Name, param.ID),
	}, nil
}
