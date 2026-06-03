package namespace

import (
	"ab/internal/dto"
	"context"
)

type namespaceService interface {
	Create(ctx context.Context, reqNamespace *dto.NameSpace) (*dto.NameSpace, error)
}
type Handler struct {
	namespaceService namespaceService
}

func New(namespaceService namespaceService) *Handler {
	return &Handler{
		namespaceService: namespaceService,
	}
}
