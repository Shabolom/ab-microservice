package customParams

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetList(ctx context.Context) ([]*dto.CustomParams, error) {
	s.logger.Info("GetList custom params started")

	params, err := s.customParamsRepo.GetList(ctx)
	if err != nil {
		s.logger.Error("GetList custom params failed",
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList custom params finished")
	return params, nil
}
