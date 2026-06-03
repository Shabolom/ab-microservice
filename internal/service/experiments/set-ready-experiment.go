package experiments

import (
	"ab/internal/dto"
	"context"

	"go.uber.org/zap"
)

func (s *Service) SetReady(ctx context.Context, experimentInfo *dto.ExperimentStatus) {
	targetExp, err := s.experimentRepo.GetExperiment(ctx, experimentInfo.ExpID)
	if err != nil {
		s.logger.Warn(
			"Failed to get experiment",
			zap.Error(err),
		)
		return
	}

	layersWithExperiments, err := s.experimentRepo.GetLayerExperimentsInPeriod(ctx, experimentInfo)
	if err != nil {
		s.logger.Info(
			"Failed to get layer with experiments",
			zap.Any("experimentInfo", experimentInfo),
			zap.Error(err))
		return
	}

	for _, layerWithExp := range layersWithExperiments {
		rolloutPercentagePerSlice := int64(0)

		for _, exp := range layerWithExp.Experiments {
			rolloutPercentagePerSlice += exp.RolloutPercentage

			if rolloutPercentagePerSlice > 100 || rolloutPercentagePerSlice < 0 {
				s.logger.Info("rollout percentage goes beyond")
				return
			}
		}
		if rolloutPercentagePerSlice+targetExp.RolloutPercentage > 100 {
			s.logger.Info("rollout percentage goes beyond")
			return
		}
	}
}
