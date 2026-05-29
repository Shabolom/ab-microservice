package split

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"fmt"

	"github.com/spaolacci/murmur3"
)

func (s *Service) Bucket(key string) int {
	hash := murmur3.Sum32([]byte(key))
	return int(hash % 100)
}

func (s *Service) InExperiment(deviceID string, experiment dto.Experiment) bool {
	key := fmt.Sprintf("experiment:%d:rollout:%s", experiment.Id, deviceID)

	bucket := s.Bucket(key)

	return bucket < experiment.RollingPercentage
}

func (s *Service) PickGroup(deviceID string, experiment dto.Experiment, groups []dto.Group) (*dto.Group, error) {
	if len(groups) == 0 {
		return nil, shortcut.ErrGroupNotFoundByBucket
	}

	if !s.InExperiment(deviceID, experiment) {
		return nil, shortcut.ErrNotInExperiment
	}

	groupKey := fmt.Sprintf("experiment:%d:group:%s", experiment.Id, deviceID)

	bucket := s.Bucket(groupKey)

	current := 0

	for i := range groups {
		current += groups[i].RollingPercentage

		if bucket < current {
			return &groups[i], nil
		}
	}

	return nil, shortcut.ErrGroupNotFoundByBucket
}
