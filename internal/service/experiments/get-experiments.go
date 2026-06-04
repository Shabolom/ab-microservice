package experiments

import (
	"ab/internal/dto"
	"ab/internal/dto/kafka-messege-dto"
	"ab/pkg/shortcut"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (s *Service) GetExperiments(parameters *dto.RequestParameters) ([]*dto.GetExperimentsReply, error) {
	s.logger.Info(
		"get experiments started",
		zap.String("namespace", parameters.NameSpace),
		zap.Int64("split_id", parameters.SplitID),
	)

	nameSpaceExperiments := s.inMemoryStorage.GetExperimentByNamespace(parameters.NameSpace)
	s.logger.Info(
		"experiments loaded from worker",
		zap.String("namespace", parameters.NameSpace),
		zap.Int("experiments_count", len(nameSpaceExperiments.RawExp)),
	)

	result := make([]*dto.GetExperimentsReply, 0)

	for _, experiment := range nameSpaceExperiments.RawExp {
		s.logger.Debug(
			"picking group for experiment",
			zap.Int64("experiment_id", experiment.Id),
			zap.String("experiment_name", experiment.Name),
			zap.Int64("rolling_percentage", experiment.RollingPercentage),
			zap.Int("groups_count", len(experiment.Group)),
		)

		group, err := s.PickGroup(parameters, &experiment)
		if errors.Is(err, shortcut.ErrGroupNotFoundByBucket) {
			s.logger.Warn(
				"group not found by bucket",
				zap.Int64("experiment_id", experiment.Id),
				zap.String("experiment_name", experiment.Name),
				zap.Error(err),
			)

			return []*dto.GetExperimentsReply{}, err
		}

		if err != nil {
			s.logger.Error(
				"failed to pick group",
				zap.Int64("experiment_id", experiment.Id),
				zap.String("experiment_name", experiment.Name),
				zap.Error(err),
			)

			return []*dto.GetExperimentsReply{}, err
		}

		if group == nil {
			s.logger.Debug(
				"experiment skipped",
				zap.Int64("experiment_id", experiment.Id),
				zap.String("experiment_name", experiment.Name),
			)

			continue
		}

		value := &dto.GetExperimentsReply{
			ExperimentName: experiment.Name,
			GroupName:      group.Name,
		}

		result = append(result, value)

		s.logger.Debug(
			"experiment matched",
			zap.Int64("experiment_id", experiment.Id),
			zap.String("experiment_name", experiment.Name),
			zap.String("group_name", group.Name),
		)
	}

	go func(result []*dto.GetExperimentsReply, userID int64) {
		for _, rawMessage := range result {
			message := &kafkaMessageDto.UserInExperimentMessage{
				UserID:         userID,
				ExperimentName: rawMessage.ExperimentName,
				GroupName:      rawMessage.GroupName,
				Timestamp:      time.Now(),
			}

			err := s.kafkaProducer.WriteEvent(message)
			if err != nil {
				s.logger.Warn(
					"failed to publish user experiment event",
					zap.Int64("user_id", userID),
					zap.String("experiment_name", rawMessage.ExperimentName),
					zap.String("group_name", rawMessage.GroupName),
					zap.Error(err),
				)

				continue
			}
		}
	}(result, parameters.SplitID)

	s.logger.Info(
		"get experiments finished",
		zap.String("namespace", parameters.NameSpace),
		zap.Int("result_count", len(result)),
	)

	return result, nil
}
