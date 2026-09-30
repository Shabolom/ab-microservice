package layer

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetList(ctx context.Context) ([]*dto.Layer, error) {
	s.logger.Info("GetList layers started")

	layers, err := s.layerRepo.GetList(ctx)
	if err != nil {
		s.logger.Warn("GetList failed",
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList layers finished")
	return layers, nil
}
