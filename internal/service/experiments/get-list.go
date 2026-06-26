package experiments

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetList(ctx context.Context) ([]*dto.Experiment, error) {
	s.logger.Info("GetList started")

	experiments, err := s.experimentRepo.GetList(ctx)
	if err != nil {
		s.logger.Warn("get list experiment err",
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList finished")
	return experiments, nil
}
