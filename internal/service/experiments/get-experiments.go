package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"errors"
)

func (s *Service) GetExperiments(ctx context.Context, splitID int64, nameSpace string) ([]*dto.GetExperimentsReply, error) {
	NameSpaceExperiments := s.inMemoryStorage.GetExperimentByNamespace(nameSpace)

	result := make([]*dto.GetExperimentsReply, 0)

	for _, experiment := range NameSpaceExperiments.RawExp {
		group, err := s.PickGroup(splitID, &experiment)
		if errors.Is(err, shortcut.ErrGroupNotFoundByBucket) {
			return []*dto.GetExperimentsReply{}, err
		}

		if group != nil {
			value := &dto.GetExperimentsReply{
				ExperimentName: experiment.Name,
				GroupName:      group.Name,
			}

			result = append(result, value)
		}
	}

	return result, nil
}
