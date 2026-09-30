package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetById(ctx context.Context, id int64) (*dto.Experiment, error) {
	s.logger.Info("GetByID started",
		zap.Int64("id", id),
	)

	if id == 0 {
		s.logger.Warn("id is 0",
			zap.Int64("id", id),
		)

		return nil, shortcut.ErrValidation
	}

	exp, err := s.experimentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	s.logger.Info("GetByID finished",
		zap.Int64("id", id),
	)
	return exp, nil
}
