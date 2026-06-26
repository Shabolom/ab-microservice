package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetByID(ctx context.Context, id int64) (*dto.NameSpace, error) {
	s.logger.Info("GetByID namespaces started")

	if id == 0 {
		s.logger.Warn("id is 0")

		return nil, shortcut.ErrValidation
	}

	namespace, err := s.namespaceRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Warn("err get by id",
			zap.Error(err),
			zap.Int64("id", id),
		)

		return nil, err
	}

	s.logger.Info("GetByID namespaces finished")
	return namespace, nil
}
