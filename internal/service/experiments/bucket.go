package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"fmt"

	"github.com/spaolacci/murmur3"
	"go.uber.org/zap"
)

func (s *Service) Bucket(key string) int64 {
	hash := murmur3.Sum32([]byte(key))
	bucket := int64(hash%100) + 1

	s.logger.Debug(
		"bucket calculated",
		zap.Int64("bucket", bucket),
	)

	return bucket
}

func (s *Service) InExperiment(splitID int64, rawExperiment *dto.RawExperiment) bool {
	key := fmt.Sprintf("experiment:%d:SplitID:%v", rawExperiment.Id, splitID)

	bucket := s.Bucket(key)
	fmt.Println(bucket, "asdasdasd")
	for _, expBucket := range rawExperiment.Bucket {
		if bucket == expBucket {
			return true
		}
	}

	s.logger.Debug(
		"bucket not in experiment",
		zap.Int64("bucket", bucket),
		zap.Int64("splitID", splitID),
	)

	return false
}

func (s *Service) PickGroup(parameters *dto.RequestParameters, rawExperiments *dto.RawExperiment) (*dto.Group, error) {
	groups := rawExperiments.Group

	for i := range groups {
		for _, id := range groups[i].DeviceID {
			if id == parameters.DeviceID {
				return &groups[i], nil
			}
		}
	}

	if rawExperiments.Status != shortcut.ExpStatusActive {
		s.logger.Debug("experiment is not active")
		return nil, nil
	}

	for _, excludedCity := range rawExperiments.ExcludedCities {
		if excludedCity == parameters.City {
			s.logger.Debug(
				"experiment filtered by excluded city",
				zap.String("city", parameters.City),
				zap.Int64("experiment_id", rawExperiments.Id),
				zap.String("experiment_name", rawExperiments.Name),
			)
			return nil, nil
		}
	}

	for _, excludedStore := range rawExperiments.ExcludedStores {
		if excludedStore == parameters.Store {
			s.logger.Debug(
				"experiment filtered by excluded store",
				zap.String("store", parameters.Store),
				zap.Int64("experiment_id", rawExperiments.Id),
				zap.String("experiment_name", rawExperiments.Name),
			)
			return nil, nil
		}
	}

	pass := false
	for _, passingCity := range rawExperiments.PassingCities {
		if passingCity == parameters.City {
			pass = true
			break
		}
	}

	if !pass {
		s.logger.Debug(
			"experiment filtered by passing cities",
			zap.String("city", parameters.City),
			zap.Int64("experiment_id", rawExperiments.Id),
			zap.String("experiment_name", rawExperiments.Name),
			zap.Strings("allowed_cities", rawExperiments.PassingCities),
		)
		return nil, nil
	}

	pass = false
	for _, passingStore := range rawExperiments.PassingStores {
		if passingStore == parameters.Store {
			pass = true
			break
		}
	}

	if !pass {
		s.logger.Debug(
			"experiment filtered by passing stores",
			zap.String("store", parameters.Store),
			zap.Int64("experiment_id", rawExperiments.Id),
			zap.String("experiment_name", rawExperiments.Name),
			zap.Strings("allowed_stores", rawExperiments.PassingStores),
		)
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

	if !s.InExperiment(parameters.SplitID, rawExperiments) {
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
