package namespace

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"
	"fmt"
)

func (h *Handler) GetNamespaceByID(ctx context.Context, req *authv1.GetNamespaceByIDRequest) (*authv1.GetNamespaceByIDReply, error) {
	namespace, err := h.namespaceService.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	return &authv1.GetNamespaceByIDReply{
		ErrInfoReason: authv1.GetNamespaceByIDReply_STATUS_OK,
		Message:       fmt.Sprintf("namespace by id %s :", namespace.Name),
		Namespace: &authv1.Namespace{
			Id:          namespace.ID,
			Name:        namespace.Name,
			Description: namespace.Description,
		},
	}, nil
}
