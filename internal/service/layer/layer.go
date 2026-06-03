package layer

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type layerRepo interface {
	Create(ctx context.Context, layer *dto.Layer) error
}

type Service struct {
	layerRepo layerRepo
	logger    *zap.Logger
}

func New(
	logger *zap.Logger,
	layerRepo layerRepo,
) *Service {
	return &Service{
		layerRepo: layerRepo,
		logger:    logger,
	}
}
