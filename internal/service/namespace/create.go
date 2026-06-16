package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, reqNamespace *dto.NameSpace) (*dto.NameSpace, error) {
	if reqNamespace.Name == "" {
		return nil, shortcut.ErrValidation
	}

	createdNamespace, err := s.namespaceRepo.Create(ctx, reqNamespace)
	if err != nil {
		s.logger.Info("Create namespace error", zap.Error(err))
		return nil, err
	}

	return createdNamespace, nil
}
