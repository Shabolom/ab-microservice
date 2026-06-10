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

	namespace, err := s.namespaceRepo.GetByID(ctx, namespaceID)
	if err != nil {
		s.logger.Warn(
			"failed to get namespace",
			zap.Int64("namespace_id", namespaceID),
			zap.Error(err),
		)

		return nil, err
	}

	if len(reqExp.ParamsGroups) > 0 {
		err = s.validateCustomGroups(ctx, reqExp, namespace.ID)
		if err != nil {
			return nil, err
		}
	}

	reqExp.Namespace = namespace.Name

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
			zap.String("namespace", reqExp.Namespace),
		)

		return nil, err
	}

	s.logger.Info(
		"experiment created",
		zap.Int64("experiment_id", createdExperiment.ID),
		zap.String("experiment_name", createdExperiment.Name),
		zap.String("namespace", createdExperiment.Namespace),
	)

	return createdExperiment, nil
}

func (s *Service) validate(reqExp *dto.Experiment) error {
	switch {
	case reqExp.Namespace == "":
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

func (s *Service) validateCustomGroups(ctx context.Context, reqExp *dto.Experiment, namespaceID int64) error {
	for _, group := range reqExp.ParamsGroups {
		if len(group.ParamsWithConditions) == 0 {
			s.logger.Warn(
				"custom param base group has no conditions",
				zap.String("experiment_name", reqExp.Name),
			)
		}
	}

	baseCustomParam, err := s.customParamRepo.GetById(ctx, reqExp.ParamsGroups[0].ParamsWithConditions[0].ParameterID)
	if err != nil {
		s.logger.Warn(
			"failed to get base custom param",
			zap.Error(err),
			zap.String("experiment_name", reqExp.Name),
			zap.Int64("parameter_id", reqExp.ParamsGroups[0].ParamsWithConditions[0].ParameterID),
		)

		return err
	}

	for groupIndex, paramGroup := range reqExp.ParamsGroups {
		if len(paramGroup.ParamsWithConditions) == 0 {
			s.logger.Warn(
				"custom param group has no conditions",
				zap.String("experiment_name", reqExp.Name),
				zap.Int("param_group_index", groupIndex),
			)

			return shortcut.ErrNoParamsInGroup
		}

		paramsIDs := make([]int64, 0, len(paramGroup.ParamsWithConditions))

		for _, paramWithCondition := range paramGroup.ParamsWithConditions {
			paramsIDs = append(paramsIDs, paramWithCondition.ParameterID)
		}

		count, err := s.customParamRepo.CountDistinctNamespaces(ctx, paramsIDs)
		if err != nil {
			s.logger.Warn(
				"failed to count distinct custom param namespaces",
				zap.Error(err),
				zap.String("experiment_name", reqExp.Name),
				zap.Int("param_group_index", groupIndex),
				zap.Int64s("parameter_ids", paramsIDs),
			)

			return err
		}

		if count != 1 {
			s.logger.Warn(
				"custom param group contains params from different namespaces",
				zap.String("experiment_name", reqExp.Name),
				zap.Int("param_group_index", groupIndex),
				zap.Int64s("parameter_ids", paramsIDs),
				zap.Int64("distinct_namespaces_count", count),
			)

			return shortcut.ErrNameSpaseCountInGroup
		}

		for _, param := range paramGroup.ParamsWithConditions {
			comparableCustomParam, err := s.customParamRepo.GetById(ctx, param.ParameterID)
			if err != nil {
				s.logger.Warn(
					"failed to get comparable custom param",
					zap.Error(err),
					zap.String("experiment_name", reqExp.Name),
					zap.Int("param_group_index", groupIndex),
					zap.Int64("parameter_id", paramGroup.ParamsWithConditions[0].ParameterID),
				)
			}

			if comparableCustomParam.NameSpaceID != baseCustomParam.NameSpaceID {
				s.logger.Warn(
					"custom param groups belong to different namespaces",
					zap.String("experiment_name", reqExp.Name),
					zap.Int("param_group_index", groupIndex),

					zap.Int64("base_parameter_id", baseCustomParam.ID),
					zap.Int64("base_namespace_id", baseCustomParam.NameSpaceID),

					zap.Int64("comparable_parameter_id", comparableCustomParam.ID),
					zap.Int64("comparable_namespace_id", comparableCustomParam.NameSpaceID),
				)

				return shortcut.ErrNamespaceIDsDontMatch
			}

			if err = shortcut.ValidateConditionValue(
				comparableCustomParam.Type,
				param.Condition,
				param.Value,
			); err != nil {
				s.logger.Warn(
					"custom parameter condition validation failed",
					zap.Int64("parameter_id", param.ParameterID),
					zap.String("parameter_type", comparableCustomParam.Type),
					zap.String("condition", param.Condition),
					zap.String("value", param.Value),
					zap.Error(err),
				)

				return err
			}
		}
	}

	if namespaceID != baseCustomParam.NameSpaceID {
		s.logger.Warn(
			"custom params namespace does not match experiment namespace",
			zap.String("experiment_name", reqExp.Name),

			zap.Int64("experiment_namespace_id", namespaceID),
			zap.Int64("custom_param_namespace_id", baseCustomParam.NameSpaceID),

			zap.Int64("base_custom_param_id", baseCustomParam.ID),
		)

		return shortcut.ErrCustomParamsNamespaceMismatch
	}

	return nil
}
