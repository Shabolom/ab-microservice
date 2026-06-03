package layer

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, reqLayer *dto.Layer) error {
	fmt.Println(reqLayer.Name, reqLayer.NameSpaceID, 123123123)
	if reqLayer.Name == "" || reqLayer.NameSpaceID <= 0 {
		s.logger.Warn(
			"layer validation failed",
			zap.String("name", reqLayer.Name),
			zap.Int64("namespace_id", reqLayer.NameSpaceID),
		)

		return shortcut.ErrValidation
	}

	err := s.layerRepo.Create(ctx, reqLayer)
	if err != nil {
		s.logger.Error(
			"failed to create layer",
			zap.String("name", reqLayer.Name),
			zap.Int64("namespace_id", reqLayer.NameSpaceID),
			zap.Error(err),
		)

		return err
	}

	return nil
}
