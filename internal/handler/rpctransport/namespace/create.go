package namespace

import (
	authv1 "ab/gen"
	"ab/internal/dto"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) CreateNamespace(ctx context.Context, req *authv1.CreateNamespaceRequest) (*authv1.StockReply, error) {
	reqNasmespace := &dto.NameSpace{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	createdNamespace, err := h.namespaceService.Create(ctx, reqNasmespace)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.StockReply{
		ErrInfoReason: authv1.StockReply_STATUS_OK,
		Message:       fmt.Sprintf("created namespace: %s, id: %v", createdNamespace.Name, createdNamespace.ID),
	}, nil
}
