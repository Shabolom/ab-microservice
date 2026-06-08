package worker

import (
	"ab/internal/dto"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) refresh(ctx context.Context) error {
	refreshCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cacheWithCustomGroups := make(map[string]dto.NameSpaceExperiments)
	cacheWithoutCustomGroups := make(map[string]dto.NameSpaceExperiments)

	expWithCustomParamGroups := make([]dto.RawExperiment, 0)
	expWithoutCustomParamGroups := make([]dto.RawExperiment, 0)

	nameSpaces, err := w.nameSpaceRepository.GetList(refreshCtx)
	if err != nil {
		return fmt.Errorf("get namespaces: %w", err)
	}

	for _, nameSpace := range nameSpaces {
		exps, err := w.experimentRepository.GetRawExperiments(refreshCtx, nameSpace.Name)
		if err != nil {
			return fmt.Errorf("get raw experiments for namespace %q: %w", nameSpace.Name, err)
		}

		for _, exp := range exps {
			if len(exp.CustomParamsGroups) > 0 {
				expWithCustomParamGroups = append(expWithCustomParamGroups, exp)
			}

			expWithoutCustomParamGroups = append(expWithoutCustomParamGroups, exp)
		}

		if len(expWithCustomParamGroups) > 0 {
			cacheWithCustomGroups[nameSpace.Name] = dto.NameSpaceExperiments{
				NameSpace: nameSpace.Name,
				RawExp:    expWithCustomParamGroups,
			}
		}

		if len(expWithoutCustomParamGroups) > 0 {
			cacheWithoutCustomGroups[nameSpace.Name] = dto.NameSpaceExperiments{
				NameSpace: nameSpace.Name,
				RawExp:    expWithoutCustomParamGroups,
			}
		}
	}

	w.rawExperimentCache.ReplaceWithCustomGroups(cacheWithCustomGroups)
	w.rawExperimentCache.ReplaceWithoutCustomGroups(cacheWithoutCustomGroups)

	w.logger.Info(
		"experiment worker refreshed",
		zap.Int("namespaces_count", len(cacheWithCustomGroups)),
	)

	return nil
}
