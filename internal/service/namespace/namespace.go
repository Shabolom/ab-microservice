package namespace

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

type namespaceRepo interface {
	Create(ctx context.Context, namespace *dto.NameSpace) (*dto.NameSpace, error)
}

type Service struct {
	namespaceRepo namespaceRepo
	logger        *zap.Logger
}

func New(
	logger *zap.Logger,
	namespaceRepo namespaceRepo,
) *Service {
	return &Service{
		namespaceRepo: namespaceRepo,
		logger:        logger,
	}
}
