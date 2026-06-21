package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"strings"

	"go.uber.org/zap"
)

func (s *Service) SetStatus(ctx context.Context, id int, status string) error {
	s.logger.Info(
		"set feature toggle status started",
		zap.Int("id", id),
		zap.String("status", status),
	)

	err := s.statusValidation(id, status)
	if err != nil {
		s.logger.Warn(
			"set feature toggle status validation failed",
			zap.Int("id", id),
			zap.String("status", status),
			zap.Error(err),
		)

		return err
	}

	err = s.featureTogglesRepo.SetStatus(ctx, id, status)
	if err != nil {
		s.logger.Error(
			"set feature toggle status failed",
			zap.Int("id", id),
			zap.String("status", status),
			zap.Error(err),
		)

		return err
	}

	s.logger.Info(
		"set feature toggle status finished",
		zap.Int("id", id),
		zap.String("status", status),
	)

	return nil
}

func (s *Service) statusValidation(id int, status string) error {
	status = strings.TrimSpace(status)
	status = strings.ToLower(status)

	switch {
	case status == "" ||
		status != dto.FeatureToggleStatusActive &&
			status != dto.FeatureToggleStatusArchived &&
			status != dto.FeatureToggleStatusDisabled &&
			status != dto.FeatureToggleStatusDraft:
		return shortcut.ErrFeatureToggleInvalidStatus

	case id == 0:
		return shortcut.ErrFeatureToggleNamespaceIDRequired
	}

	return nil
}
