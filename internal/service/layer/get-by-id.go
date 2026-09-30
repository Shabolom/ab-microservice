package layer

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetById(ctx context.Context, id int64) (*dto.Layer, error) {
	s.logger.Info("GetById layer started",
		zap.Int64("id", id),
	)

	if id == 0 {
		s.logger.Info("id is 0")
		return nil, shortcut.ErrValidation
	}

	layer, err := s.layerRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("GetById layer failed",
			zap.Int64("id", id),
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList layers started",
		zap.Int64("id", id),
	)

	return layer, nil
}
