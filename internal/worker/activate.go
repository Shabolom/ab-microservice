package worker

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) startActivateExperimentWorker(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	if err := w.activateExperimentWorker(ctx); err != nil {
		w.logger.Error("activate experiment iteration failed", zap.Error(err))
	}

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("activate experiment worker stopped", zap.Error(ctx.Err()))
			return

		case <-ticker.C:
			if err := w.activateExperimentWorker(ctx); err != nil {
				w.logger.Error("activate experiment iteration failed", zap.Error(err))
			}
		}
	}
}

func (w *Worker) activateExperimentWorker(ctx context.Context) error {

	layersWithExperiments, err := w.experimentRepository.GetLayersWithExperiments(ctx)
	if err != nil {
		return fmt.Errorf("get layers with experiments: %w", err)
	}

	usedBuckets := make([]int64, 0, 100)
	readyExps := make([]*dto.Experiment, 0)

	for _, layer := range layersWithExperiments {
		for i := range layer.Experiments {
			exp := &layer.Experiments[i]

			if exp.Status == shortcut.ExpStatusActive {
				usedBuckets = append(usedBuckets, exp.Bucket...)
				continue
			}

			if exp.Status == shortcut.ExpStatusReady {
				readyExps = append(readyExps, exp)
			}
		}
	}

	for _, exp := range readyExps {
		buckets, err := generateBuckets(usedBuckets, exp.RolloutPercentage)
		if err != nil {
			w.logger.Error(
				"generate buckets failed",
				zap.Int64("experiment_id", exp.ID),
				zap.Int64("rollout_percentage", exp.RolloutPercentage),
				zap.Error(err),
			)

			continue
		}

		err = w.experimentRepository.UpdateStatusBuckets(
			ctx,
			exp.ID,
			shortcut.ExpStatusActive,
			buckets,
		)
		if err != nil {
			w.logger.Error(
				"update status buckets failed",
				zap.Int64("experiment_id", exp.ID),
				zap.Int64s("buckets", buckets),
				zap.Error(err),
			)

			continue
		}

		usedBuckets = append(usedBuckets, buckets...)

		w.logger.Info(
			"experiment activated",
			zap.Int64("experiment_id", exp.ID),
			zap.Int64("rollout_percentage", exp.RolloutPercentage),
			zap.Int64s("buckets", buckets),
		)
	}

	return nil
}

func generateBuckets(usedBuckets []int64, count int64) ([]int64, error) {
	used := make(map[int64]struct{}, len(usedBuckets))

	for _, bucket := range usedBuckets {
		if bucket < 1 || bucket > 100 {
			return nil, fmt.Errorf("invalid bucket: %d", bucket)
		}

		used[bucket] = struct{}{}
	}

	freeBuckets := make([]int64, 0, 100-len(used))

	for bucket := int64(1); bucket <= 100; bucket++ {
		if _, ok := used[bucket]; !ok {
			freeBuckets = append(freeBuckets, bucket)
		}
	}

	if int64(len(freeBuckets)) < count {
		return nil, fmt.Errorf("not enough free buckets")
	}

	rand.Shuffle(len(freeBuckets), func(i, j int) {
		freeBuckets[i], freeBuckets[j] = freeBuckets[j], freeBuckets[i]
	})

	return freeBuckets[:count], nil
}
