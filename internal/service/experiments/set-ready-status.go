package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"math/rand"
	"slices"

	"go.uber.org/zap"
)

func (s *Service) SetReady(ctx context.Context, targetExpID int64) error {
	if targetExpID == 0 {
		s.logger.Warn("empty experiment id")
		return shortcut.ErrValidation
	}

	s.logger.Info(
		"set experiment ready started",
		zap.Int64("experiment_id", targetExpID),
	)

	targetExp, err := s.experimentRepo.GetExperimentWithLayers(ctx, targetExpID)
	if err != nil {
		s.logger.Warn(
			"failed to get experiment",
			zap.Int64("experiment_id", targetExpID),
			zap.Error(err),
		)

		return shortcut.ErrFailedToGetExperiment
	}

	if targetExp.Status == shortcut.ExpStatusActive || targetExp.Status == shortcut.ExpStatusReady {
		s.logger.Warn(
			"experiment already active or ready",
			zap.Int64("experiment_id", targetExp.ID),
			zap.String("status", targetExp.Status),
		)

		return shortcut.ErrExperimentAlreadyRunning
	}

	layersBuckets, err := s.experimentRepo.GetLayerBucketsInPeriod(ctx, targetExp)
	if err != nil {
		s.logger.Warn(
			"failed to get occupied buckets in experiment layers",
			zap.Int64("experiment_id", targetExp.ID),
			zap.Time("started_at", targetExp.StartDate),
			zap.Time("ended_at", targetExp.EndDate),
			zap.Error(err),
		)

		return shortcut.ErrFailedToGetLayerExperiments
	}

	if len(layersBuckets) == 0 {
		layersBuckets, err = s.createLayerExp(ctx, targetExp.Namespace, targetExp.ID)
		if err != nil {
			s.logger.Error(
				"create layer experiments failed",
				zap.Int64("experiment_id", targetExp.ID),
				zap.String("namespace", targetExp.Namespace),
				zap.Error(err),
			)

			return err
		}
	}

	s.logger.Info(
		"loaded layer buckets",
		zap.Int64("experiment_id", targetExp.ID),
		zap.Int("layers_count", len(layersBuckets)),
	)

	newLayerBuckets := make([]dto.LayerBuckets, 0, len(layersBuckets))

	for _, layerBucket := range layersBuckets {
		usedBuckets := len(layerBucket.Buckets)
		availableBuckets := 100 - usedBuckets

		s.logger.Debug(
			"layer buckets analysis",
			zap.Int64("experiment_id", targetExp.ID),
			zap.Int64("layer_id", layerBucket.LayerID),
			zap.Int("used_buckets", usedBuckets),
			zap.Int("available_buckets", availableBuckets),
			zap.Int64("required_buckets", targetExp.RolloutPercentage),
		)

		if int64(usedBuckets)+targetExp.RolloutPercentage > 100 {
			s.logger.Warn(
				"not enough free buckets in layer for experiment rollout",
				zap.Int64("experiment_id", targetExp.ID),
				zap.Int64("layer_id", layerBucket.LayerID),
				zap.Int64("required_buckets", targetExp.RolloutPercentage),
				zap.Int("used_buckets", usedBuckets),
				zap.Int("available_buckets", availableBuckets),
				zap.Time("started_at", targetExp.StartDate),
				zap.Time("ended_at", targetExp.EndDate),
			)

			return shortcut.ErrExperimentLayerRolloutTooBig
		}

		generatedBuckets := s.generateBuckets(
			layerBucket.Buckets,
			targetExp.RolloutPercentage,
		)

		s.logger.Debug(
			"generated buckets for layer",
			zap.Int64("experiment_id", targetExp.ID),
			zap.Int64("layer_id", layerBucket.LayerID),
			zap.Int("generated_count", len(generatedBuckets)),
		)

		newLayerBuckets = append(newLayerBuckets, dto.LayerBuckets{
			LayerID: layerBucket.LayerID,
			Buckets: generatedBuckets,
		})
	}

	err = s.experimentRepo.SetReadyStatus(
		ctx,
		targetExp.ID,
		shortcut.ExpStatusReady,
		newLayerBuckets,
	)
	if err != nil {
		s.logger.Warn(
			"failed to update experiment status and layer buckets",
			zap.Int64("experiment_id", targetExp.ID),
			zap.String("status", shortcut.ExpStatusReady),
			zap.Int("layers_count", len(newLayerBuckets)),
			zap.Error(err),
		)

		return shortcut.ErrFailedToUpdateExperiment
	}

	s.logger.Info(
		"experiment marked as ready",
		zap.Int64("experiment_id", targetExp.ID),
		zap.Int("layers_count", len(newLayerBuckets)),
		zap.Int64("rollout_percentage", targetExp.RolloutPercentage),
	)

	return nil
}

func (s *Service) generateBuckets(usedBuckets []int64, count int64) []int64 {
	used := make(map[int64]struct{}, len(usedBuckets))

	for _, bucket := range usedBuckets {
		used[bucket] = struct{}{}
	}

	freeBuckets := make([]int64, 0, 100-len(used))

	for bucket := int64(1); bucket <= 100; bucket++ {
		if _, ok := used[bucket]; !ok {
			freeBuckets = append(freeBuckets, bucket)
		}
	}

	rand.Shuffle(len(freeBuckets), func(i, j int) {
		freeBuckets[i], freeBuckets[j] = freeBuckets[j], freeBuckets[i]
	})

	if count > int64(len(freeBuckets)) {
		count = int64(len(freeBuckets))
	}

	result := freeBuckets[:count]

	slices.Sort(result)

	return result
}

func (s *Service) createLayerExp(ctx context.Context, namespaceName string, expID int64) ([]dto.LayerBuckets, error) {
	namespace, err := s.namespaceRepo.GetByName(ctx, namespaceName)
	if err != nil {
		return nil, err
	}

	layersID, err := s.layerRepo.GetIDsByNamespaceId(ctx, namespace.ID)
	if err != nil {
		return nil, err
	}

	layersBuckets, err := s.experimentRepo.CreateAndGetLayerExperiment(ctx, layersID, expID)
	if err != nil {
		return nil, err
	}

	return layersBuckets, nil
}
