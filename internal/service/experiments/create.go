package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"time"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, reqExp *dto.Experiment) (*dto.Experiment, error) {
	reqExp.Status = shortcut.ExpStatusInTest

	if err := s.validate(reqExp); err != nil {
		s.logger.Warn(
			"experiment validation failed",
			zap.Error(err),
			zap.String("experiment_name", reqExp.Name),
		)

		return nil, err
	}

	createdExperiment, err := s.experimentRepo.CreateExperiment(ctx, reqExp)
	if err != nil {
		s.logger.Error(
			"create experiment failed",
			zap.Error(err),
		)

		return nil, err
	}

	return createdExperiment, nil
}

func (s *Service) validate(reqExp *dto.Experiment) error {
	switch {
	case reqExp.EndDate.Before(time.Now()):
		return shortcut.ErrExperimentEndDateInPast

	case !reqExp.StartDate.Before(reqExp.EndDate):
		return shortcut.ErrExperimentStartDateAfterEnd

	case reqExp.Name == "":
		return shortcut.ErrExperimentNameRequired

	case len(reqExp.Groups) < 2:
		return shortcut.ErrExperimentGroupsMinCount

	case len(reqExp.LayersID) == 0:
		return shortcut.ErrExperimentLayersRequired

	case reqExp.RolloutPercentage > 100:
		return shortcut.ErrExperimentRolloutOutOfRange
	}

	groupRolloutPercentage := int64(0)

	for _, group := range reqExp.Groups {
		groupRolloutPercentage += group.RollingPercentage

		if groupRolloutPercentage > 100 {
			return shortcut.ErrExperimentGroupsRolloutTooBig
		}
	}

	return nil
}
