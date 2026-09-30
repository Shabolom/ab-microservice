package layer

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetLayers(ctx context.Context, req *emptypb.Empty) (*authv1.GetLayerBysReply, error) {
	layers, err := h.layerService.GetList(ctx)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var layersResponse []*authv1.Layer
	for _, layer := range layers {
		layersResponse = append(layersResponse, &authv1.Layer{
			Id:          layer.ID,
			NamespaceId: layer.NameSpaceID,
			Name:        layer.Name,
			Description: layer.Description,
		})
	}

	return &authv1.GetLayerBysReply{
		ErrInfoReason: authv1.GetLayerBysReply_STATUS_OK,
		Message:       "layers:",
		Layers:        layersResponse,
	}, nil
}
