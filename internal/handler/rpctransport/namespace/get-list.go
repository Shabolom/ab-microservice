package namespace

import (
	authv1 "ab/gen"
	"ab/internal/render"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetNamespaces(ctx context.Context, req *emptypb.Empty) (*authv1.GetNamespaceBysReply, error) {
	namespaces, err := h.namespaceService.GetList(ctx)
	if err != nil {
		return nil, render.ErrorValidator(err)
	}

	var responseNamespaces []*authv1.Namespace
	for _, namespace := range namespaces {
		responseNamespaces = append(responseNamespaces, &authv1.Namespace{
			Id:          namespace.ID,
			Name:        namespace.Name,
			Description: namespace.Description,
		})
	}

	return &authv1.GetNamespaceBysReply{
		ErrInfoReason: authv1.GetNamespaceBysReply_STATUS_OK,
		Message:       "namespaces:",
		Namespaces:    responseNamespaces,
	}, nil
}
