package experiments

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

func (s *Service) InExperiment(deviceID int64, rawExperiment *dto.RawExperiment) bool {
	key := fmt.Sprintf("experiment:%d:device-id:%s", rawExperiment.Id, deviceID)

	bucket := s.Bucket(key)

	return bucket < rawExperiment.RollingPercentage
}

func (s *Service) PickGroup(deviceID int64, rawExperiments *dto.RawExperiment) (*dto.Group, error) {
	groups := rawExperiments.Group

	if len(groups) == 0 {
		return nil, shortcut.ErrGroupNotFoundByBucket
	}

	if !s.InExperiment(deviceID, rawExperiments) {
		return nil, nil
	}

	groupKey := fmt.Sprintf("experiment:%d:device-id:%s", rawExperiments.Id, deviceID)

	bucket := s.Bucket(groupKey)

	current := 0

	for i := range groups {
		current += groups[i].RollingPercentage

		if bucket <= current {
			return &groups[i], nil
		}
	}

	return nil, nil
}
