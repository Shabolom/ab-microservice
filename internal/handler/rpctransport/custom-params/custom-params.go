package customParams

import (
	"ab/internal/dto"
	"context"
)

type customParamsService interface {
	Create(ctx context.Context, param *dto.CustomParams) (*dto.CustomParams, error)
	GetById(ctx context.Context, id int64) (*dto.CustomParams, error)
	GetList(ctx context.Context) ([]*dto.CustomParams, error)
}
type Handler struct {
	customParamsService customParamsService
}

func New(customParamsService customParamsService) *Handler {
	return &Handler{
		customParamsService: customParamsService,
	}
}
