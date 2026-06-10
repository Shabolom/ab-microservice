package worker

import (
	"ab/internal/dto"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) refresh(ctx context.Context) error {
	refreshCtx, cancel := context.WithTimeout(ctx, time.Duration(w.ctxInterval)*time.Second)
	defer cancel()

	cacheWithCustomGroups := make(map[string]dto.NameSpaceExperiments)
	cacheWithoutCustomGroups := make(map[string]dto.NameSpaceExperiments)

	namespaces, err := w.namespaceRepository.GetList(refreshCtx)
	if err != nil {
		return fmt.Errorf("get namespaces: %w", err)
	}

	for _, namespace := range namespaces {
		expWithCustomParamGroups := make([]dto.RawExperiment, 0)
		expWithoutCustomParamGroups := make([]dto.RawExperiment, 0)

		exps, err := w.experimentRepository.GetRawExperiments(refreshCtx, namespace.Name)
		if err != nil {
			return fmt.Errorf("get raw experiments for namespace %q: %w", namespace.Name, err)
		}

		for _, exp := range exps {
			if len(exp.CustomParamsGroups) > 0 {
				expWithCustomParamGroups = append(expWithCustomParamGroups, exp)
				continue
			} else {
				expWithoutCustomParamGroups = append(expWithoutCustomParamGroups, exp)
			}
		}

		if len(expWithCustomParamGroups) > 0 {
			cacheWithCustomGroups[namespace.Name] = dto.NameSpaceExperiments{
				NameSpace: namespace.Name,
				RawExp:    expWithCustomParamGroups,
			}
		}

		if len(expWithoutCustomParamGroups) > 0 {
			cacheWithoutCustomGroups[namespace.Name] = dto.NameSpaceExperiments{
				NameSpace: namespace.Name,
				RawExp:    expWithoutCustomParamGroups,
			}
		}
	}

	w.rawExperimentCache.ReplaceWithCustomGroups(cacheWithCustomGroups)
	w.rawExperimentCache.ReplaceWithoutCustomGroups(cacheWithoutCustomGroups)

	w.logger.Info(
		"experiment worker refreshed",
		zap.Int("cacheWithCustomGroups_count", len(cacheWithCustomGroups)),
		zap.Int("cacheWithoutCustomGroups_count", len(cacheWithoutCustomGroups)),
	)

	return nil
}
