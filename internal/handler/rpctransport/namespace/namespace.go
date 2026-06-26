package namespace

import (
	"ab/internal/dto"
	"context"
)

type namespaceService interface {
	Create(ctx context.Context, reqNamespace *dto.NameSpace) (*dto.NameSpace, error)
	GetByID(ctx context.Context, id int64) (*dto.NameSpace, error)
	GetList(ctx context.Context) ([]*dto.NameSpace, error)
}
type Handler struct {
	namespaceService namespaceService
}

func New(namespaceService namespaceService) *Handler {
	return &Handler{
		namespaceService: namespaceService,
	}
}
