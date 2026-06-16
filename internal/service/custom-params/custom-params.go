package customParams

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type customParamsRepo interface {
	Create(ctx context.Context, param *dto.CustomParams) (*dto.CustomParams, error)
	GetById(ctx context.Context, id int64) (dto.CustomParams, error)
}

type Service struct {
	customParamsRepo customParamsRepo
	logger           *zap.Logger
}

func New(
	logger *zap.Logger,
	customParamsRepo customParamsRepo,
) *Service {
	return &Service{
		customParamsRepo: customParamsRepo,
		logger:           logger,
	}
}
