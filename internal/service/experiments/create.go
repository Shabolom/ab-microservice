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

	s.logger.Info(
		"create experiment started",
		zap.String("experiment_name", reqExp.Name),
		zap.Int64s("layer_ids", reqExp.LayersID),
	)

	layers, err := s.layerRepo.GetLayers(ctx, reqExp.LayersID)
	if err != nil {
		s.logger.Warn(
			"experiment layers error",
			zap.Error(err),
			zap.String("experiment_name", reqExp.Name),
		)

		return nil, err
	}

	if len(layers) == 0 {
		s.logger.Warn(
			"no layers found for experiment",
			zap.String("experiment_name", reqExp.Name),
		)

		return nil, shortcut.ErrValidation
	}

	namespaceID := layers[0].NameSpaceID
	for _, layer := range layers {
		if layer.NameSpaceID != namespaceID {
			s.logger.Warn(
				"layers belong to different namespaces",
				zap.String("experiment_name", reqExp.Name),
				zap.Int64("expected_namespace_id", namespaceID),
				zap.Int64("actual_namespace_id", layer.NameSpaceID),
				zap.Int64("layer_id", layer.ID),
			)

			return nil, shortcut.ErrDifferentNamespaces
		}
	}

	namespace, err := s.nameSpaceRepo.GetByID(ctx, namespaceID)
	if err != nil {
		s.logger.Warn(
			"failed to get namespace",
			zap.Int64("namespace_id", namespaceID),
			zap.Error(err),
		)

		return nil, err
	}

	reqExp.NameSpace = namespace.Name

	if err = s.validate(reqExp); err != nil {
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
			zap.String("experiment_name", reqExp.Name),
			zap.String("namespace", reqExp.NameSpace),
		)

		return nil, err
	}

	s.logger.Info(
		"experiment created",
		zap.Int64("experiment_id", createdExperiment.ID),
		zap.String("experiment_name", createdExperiment.Name),
		zap.String("namespace", createdExperiment.NameSpace),
	)

	return createdExperiment, nil
}

func (s *Service) validate(reqExp *dto.Experiment) error {
	switch {
	case reqExp.NameSpace == "":
		s.logger.Warn("experiment namespace is empty")
		return shortcut.ErrValidation

	case reqExp.EndDate.Before(time.Now()):
		s.logger.Warn(
			"experiment end date is in past",
			zap.Time("end_date", reqExp.EndDate),
		)
		return shortcut.ErrExperimentEndDateInPast

	case !reqExp.StartDate.Before(reqExp.EndDate):
		s.logger.Warn(
			"experiment start date is after end date",
			zap.Time("start_date", reqExp.StartDate),
			zap.Time("end_date", reqExp.EndDate),
		)
		return shortcut.ErrExperimentStartDateAfterEnd

	case reqExp.Name == "":
		s.logger.Warn("experiment name is empty")
		return shortcut.ErrExperimentNameRequired

	case len(reqExp.Groups) < 2:
		s.logger.Warn(
			"experiment has less than 2 groups",
			zap.Int("groups_count", len(reqExp.Groups)),
		)
		return shortcut.ErrExperimentGroupsMinCount

	case len(reqExp.LayersID) == 0:
		s.logger.Warn("experiment has no layers")
		return shortcut.ErrExperimentLayersRequired

	case reqExp.RolloutPercentage > 100:
		s.logger.Warn(
			"experiment rollout percentage is too large",
			zap.Int64("rollout_percentage", reqExp.RolloutPercentage),
		)
		return shortcut.ErrExperimentRolloutOutOfRange
	}

	groupRolloutPercentage := int64(0)

	for _, group := range reqExp.Groups {
		groupRolloutPercentage += group.RollingPercentage

		if groupRolloutPercentage > 100 {
			s.logger.Warn(
				"groups rollout percentage exceeds 100",
				zap.Int64("current_rollout_percentage", groupRolloutPercentage),
			)

			return shortcut.ErrExperimentGroupsRolloutTooBig
		}
	}

	if groupRolloutPercentage != 100 {
		s.logger.Warn(
			"groups rollout percentage is not equal to 100",
			zap.Int64("current_rollout_percentage", groupRolloutPercentage),
		)

		return shortcut.ErrExperimentGroupsRolloutNotFull
	}

	return nil
}
