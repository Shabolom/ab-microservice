package experiments

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"math/rand"

	"go.uber.org/zap"
)

func (s *Service) SetReady(ctx context.Context, targetExpID int64) error {
	if targetExpID == 0 {
		return shortcut.ErrValidation
	}

	targetExp, err := s.experimentRepo.GetExperiment(ctx, targetExpID)
	if err != nil {
		s.logger.Warn(
			"failed to get experiment",
			zap.Int64("experiment_id", targetExpID),
			zap.Error(err),
		)

		return shortcut.ErrFailedToGetExperiment
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

	newLayerBuckets := make([]dto.LayerBuckets, 0, len(layersBuckets))

	for _, layerBucket := range layersBuckets {
		usedBuckets := len(layerBucket.Buckets)
		availableBuckets := 100 - usedBuckets

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

		newLayerBuckets = append(newLayerBuckets, dto.LayerBuckets{
			LayerID: layerBucket.LayerID,
			Buckets: generatedBuckets,
		})
	}

	err = s.experimentRepo.UpdateStatusBucketsTx(
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

	return nil
}

//func (s *Service) validateRolloutPercentage(
//	layersWithExp []dto.LayerWithExperiments,
//	targetExpRolloutPercentage int64,
//) (map[int64][]int64, error) {
//	usedBucketsInlayer := make(map[int64][]int64, 100)
//
//	for _, layerWithExp := range layersWithExp {
//		usedRolloutPercentageInLayer := int64(0)
//		val, ok := usedBucketsInlayer[layerWithExp.LayerID]
//		if !ok {
//			usedBucketsInlayer[layerWithExp.LayerID] = make([]int64, 0, 100)
//		}
//
//		for _, exp := range layerWithExp.Experiments {
//			usedRolloutPercentageInLayer += exp.RolloutPercentage
//
//			for _, bucket := range exp.Bucket {
//				val = append(val, bucket)
//			}
//		}
//
//		if usedRolloutPercentageInLayer+targetExpRolloutPercentage > 100 {
//			s.logger.Info(
//				"rollout percentage goes beyond limit",
//				zap.Int64("layer_id", layerWithExp.LayerID),
//				zap.Int64("used_rollout_percentage_in_layer", usedRolloutPercentageInLayer),
//				zap.Int64("target_rollout_percentage", targetExpRolloutPercentage),
//			)
//
//			return nil, shortcut.ErrExperimentLayerRolloutTooBig
//		}
//	}
//
//	return usedBucketsInlayer, nil
//}

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

	return freeBuckets[:count]
}
