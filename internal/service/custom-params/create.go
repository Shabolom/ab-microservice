package customParams

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"strings"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, param *dto.CustomParams) (*dto.CustomParams, error) {
	if !s.paramValidation(param.Type) {
		return nil, shortcut.ErrValidation
	}

	result, err := s.customParamsRepo.Create(ctx, param)
	if err != nil {
		s.logger.Warn(
			"create custom parameter failed",
			zap.String("name", param.Name),
			zap.String("type", param.Type),
			zap.Error(err),
		)
		return nil, err
	}

	return result, nil
}

func (s *Service) paramValidation(paramType string) bool {
	paramType = strings.ToUpper(paramType)

	switch paramType {
	case shortcut.ParamTypeInt:
		return true
	case shortcut.ParamTypeDate:
		return true
	case shortcut.ParamTypeString:
		return true
	case shortcut.ParamTypeSemver:
		return true
	default:
		return false
	}
}
