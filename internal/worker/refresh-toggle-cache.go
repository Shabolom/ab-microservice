package worker

import (
	"ab/internal/dto"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func (w *Worker) refreshToggleCache(ctx context.Context) error {
	refreshCtx, cancel := context.WithTimeout(
		ctx,
		time.Duration(w.ctxInterval)*time.Second,
	)
	defer cancel()

	namespaces, err := w.namespaceRepository.GetList(refreshCtx)
	if err != nil {
		w.logger.Error(
			"error getting namespaces",
			zap.Error(err),
		)

		return fmt.Errorf("get namespaces: %w", err)
	}

	cacheFeatureToggles := make(map[string]dto.NamespaceFeatureToggle)

	for _, namespace := range namespaces {
		featureToggles, err := w.featureToggleRepo.GetActiveByNamespaceID(
			refreshCtx,
			namespace.ID,
		)
		if err != nil {
			w.logger.Error(
				"error getting feature toggles",
				zap.String("namespace", namespace.Name),
				zap.Int64("namespace_id", namespace.ID),
				zap.Error(err),
			)

			return fmt.Errorf("get feature toggles: %w", err)
		}

		cacheFeatureToggles[namespace.Name] = dto.NamespaceFeatureToggle{
			Namespace:      namespace.Name,
			FeatureToggles: featureToggles,
		}
	}

	w.rawExperimentCache.ReplaceFeatureToggle(cacheFeatureToggles)

	return nil
}
