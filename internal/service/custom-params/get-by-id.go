package customParams

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetById(ctx context.Context, id int64) (*dto.CustomParams, error) {
	s.logger.Info("GetById custom params started")

	if id == 0 {
		s.logger.Info("id is 0")
		return nil, shortcut.ErrValidation
	}

	customParam, err := s.customParamsRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Warn("err GetById custom param",
			zap.Error(err),
			zap.Int64("id", id),
		)

		return nil, err
	}

	s.logger.Info("GetById custom params finished")
	return customParam, nil
}
