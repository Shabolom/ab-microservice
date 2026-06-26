package namespace

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) GetList(ctx context.Context) ([]*dto.NameSpace, error) {
	s.logger.Info("GetList namespaces Started")

	namespaces, err := s.namespaceRepo.GetList(ctx)
	if err != nil {
		s.logger.Warn("GetList Failed",
			zap.Error(err),
		)

		return nil, err
	}

	s.logger.Info("GetList namespaces Finished")
	return namespaces, nil
}
