package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"fmt"
	"time"

	"github.com/spaolacci/murmur3"
	"go.uber.org/zap"
)

func (s *Service) Bucket(key string) int64 {
	hash := murmur3.Sum32([]byte(key))
	bucket := int64(hash%100) + 1

	s.logger.Debug(
		"bucket calculated",
		zap.String("key", key),
		zap.Uint32("hash", hash),
		zap.Int64("bucket", bucket),
	)

	return bucket
}

func (s *Service) InExperiment(splitID int64, rollingPercentageExp int64, rawExperiment *dto.RawExperiment) bool {
	key := fmt.Sprintf("experiment:%d:SplitID:%v", rawExperiment.Id, splitID)

	bucket := s.Bucket(key)
	inExperiment := bucket <= rawExperiment.RollingPercentage+rollingPercentageExp

	s.logger.Debug(
		"in experiment checked",
		zap.Int64("experiment_id", rawExperiment.Id),
		zap.Int64("split_id", splitID),
		zap.Int64("bucket", bucket),
		zap.Int64("rolling_percentage", rawExperiment.RollingPercentage),
		zap.Bool("in_experiment", inExperiment),
	)

	return inExperiment
}

func (s *Service) PickGroup(parameters *dto.RequestParameters, rawExperiments *dto.RawExperiment, rollingPercentageExp int64) (*dto.Group, error) {
	groups := rawExperiments.Group
	for i := range groups {
		for _, id := range groups[i].DeviceID {
			if id == parameters.DeviceID {
				return &groups[i], nil
			}
		}
	}

	if rawExperiments.Status != "active" || rawExperiments.EndDate.After(time.Now()) || rawExperiments.StartDate.Before(time.Now()) {
		return nil, nil
	}

	s.logger.Debug(
		"pick group started",
		zap.Int64("experiment_id", rawExperiments.Id),
		zap.Int64("split_id", parameters.SplitID),
		zap.Int64("device_id", parameters.DeviceID),
		zap.Int("groups_count", len(groups)),
	)

	if len(groups) == 0 {
		s.logger.Warn(
			"experiment has no groups",
			zap.Int64("experiment_id", rawExperiments.Id),
		)

		return nil, shortcut.ErrGroupNotFoundByBucket
	}

	if !s.InExperiment(parameters.SplitID, rollingPercentageExp, rawExperiments) {
		s.logger.Debug(
			"split is not in experiment",
			zap.Int64("experiment_id", rawExperiments.Id),
			zap.Int64("split_id", parameters.SplitID),
		)

		return nil, nil
	}

	groupKey := fmt.Sprintf("rawExperiments-id:%d:SplitID:%v", rawExperiments.Id, parameters.SplitID)

	bucket := s.Bucket(groupKey)
	current := int64(0)

	for i := range groups {
		current += groups[i].RollingPercentage

		s.logger.Debug(
			"group bucket range checked",
			zap.Int64("experiment_id", rawExperiments.Id),
			zap.String("group_name", groups[i].Name),
			zap.Int64("group_rolling_percentage", groups[i].RollingPercentage),
			zap.Int64("current", current),
			zap.Int64("bucket", bucket),
		)

		if bucket <= current {
			s.logger.Debug(
				"group picked",
				zap.Int64("experiment_id", rawExperiments.Id),
				zap.String("group_name", groups[i].Name),
				zap.Int64("bucket", bucket),
			)

			return &groups[i], nil
		}
	}

	s.logger.Info(
		"group was not picked",
		zap.Int64("experiment_id", rawExperiments.Id),
		zap.Int64("bucket", bucket),
		zap.Int64("total_groups_percentage", current),
	)

	return nil, nil
}
