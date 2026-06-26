package layer

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) GetLayerByID(ctx context.Context, req *authv1.GetLayerByIDRequest) (*authv1.GetLayerByIDReply, error) {
	id := req.GetId()

	layers, err := h.layerService.GetById(ctx, id)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.GetLayerByIDReply{
		ErrInfoReason: authv1.GetLayerByIDReply_STATUS_OK,
		Message:       fmt.Sprintf("layer by id %v", id),
		Layer: &authv1.Layer{
			Id:          layers.ID,
			NamespaceId: layers.NameSpaceID,
			Name:        layers.Name,
			Description: layers.Description,
		},
	}, nil
}
