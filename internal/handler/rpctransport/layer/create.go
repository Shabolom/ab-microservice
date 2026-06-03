package layer

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) CreateLayer(ctx context.Context, req *authv1.CreateLayerRequest) (*authv1.StockReply, error) {
	reqLayer := &dto.Layer{
		NameSpaceID: req.GetNameSpaceID(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	err := h.layerService.Create(ctx, reqLayer)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("layer %v created id: %v", reqLayer.Name, reqLayer.ID),
	}, nil
}
