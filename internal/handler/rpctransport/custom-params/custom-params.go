package customParams

import (
	"ab/internal/dto"
	"context"
)

type customParamsService interface {
	Create(ctx context.Context, param *dto.CustomParams) (*dto.CustomParams, error)
}
type Handler struct {
	customParamsService customParamsService
}

func New(customParamsService customParamsService) *Handler {
	return &Handler{
		customParamsService: customParamsService,
	}
}
