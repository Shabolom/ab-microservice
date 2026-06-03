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

	cache := make(map[string]dto.NameSpaceExperiments)

	nameSpaces, err := w.nameSpaceRepository.GetList(refreshCtx)
	if err != nil {
		return fmt.Errorf("get namespaces: %w", err)
	}

	for _, nameSpace := range nameSpaces {
		exp, err := w.experimentRepository.GetRawExperiments(refreshCtx, nameSpace.Name)
		if err != nil {
			return fmt.Errorf("get raw experiments for namespace %q: %w", nameSpace.Name, err)
		}

		cache[nameSpace.Name] = dto.NameSpaceExperiments{
			NameSpace: nameSpace.Name,
			RawExp:    exp,
		}
	}

	w.rawExperimentCache.Replace(cache)

	w.logger.Info(
		"experiment worker refreshed",
		zap.Int("namespaces_count", len(cache)),
	)

	return nil
}
