package layer

import (
	"ab/internal/dto"
	"context"
)

type layerService interface {
	Create(ctx context.Context, reqLayer *dto.Layer) error
	GetById(ctx context.Context, id int64) (*dto.Layer, error)
	GetList(ctx context.Context) ([]*dto.Layer, error)
}
type Handler struct {
	layerService layerService
}

func New(layerService layerService) *Handler {
	return &Handler{
		layerService: layerService,
	}
}
