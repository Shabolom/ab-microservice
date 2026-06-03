package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"
)

func (s *Storage) GetLayerExperimentsInPeriod(ctx context.Context, experimentInfo *dto.ExperimentStatus) ([]dto.LayerWithExperiments, error) {
	query := `
		SELECT
			le.layer_id,
			e.id,
			e.name,
			e.status,
			e.rollout_percentage,
			e.start_date,
			e.end_date,
			e.bucket,
			COALESCE(
				array_agg(DISTINCT all_le.layer_id),
				'{}'
			) AS layers_id
		FROM experiments e
		JOIN layer_experiments le
			ON le.experiment_id = e.id
		LEFT JOIN layer_experiments all_le
			ON all_le.experiment_id = e.id
		WHERE le.layer_id = $1
		  AND e.status IN ($2, $3)
		  AND e.start_date <= $4
		  AND e.end_date >= $5
		  AND e.id <> $6
		GROUP BY
			le.layer_id,
			e.id,
			e.name,
			e.status,
			e.rollout_percentage,
			e.start_date,
			e.end_date,
			e.bucket
		ORDER BY le.layer_id, e.id
	`

	rows, err := s.conn.Query(
		ctx,
		query,
		experimentInfo.LayerID,
		shortcut.ExpStatusReady,
		shortcut.ExpStatusActive,
		experimentInfo.EndedAt,
		experimentInfo.StartedAt,
		experimentInfo.ExpID,
	)
	if err != nil {
		return nil, fmt.Errorf("get layer experiments in period: %w", err)
	}
	defer rows.Close()

	layersMap := make(map[int64]*dto.LayerWithExperiments)

	for rows.Next() {
		var (
			layerID int64
			exp     dto.Experiment
		)

		err = rows.Scan(
			&layerID,
			&exp.ID,
			&exp.Name,
			&exp.Status,
			&exp.RolloutPercentage,
			&exp.StartDate,
			&exp.EndDate,
			&exp.Bucket,
			&exp.LayersID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan layer experiment: %w", err)
		}

		layer, ok := layersMap[layerID]
		if !ok {
			layer = &dto.LayerWithExperiments{
				LayerID:     layerID,
				Experiments: make([]dto.Experiment, 0),
			}
			layersMap[layerID] = layer
		}

		layer.Experiments = append(layer.Experiments, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate layer experiments: %w", err)
	}

	result := make([]dto.LayerWithExperiments, 0, len(layersMap))

	for _, layer := range layersMap {
		result = append(result, *layer)
	}

	return result, nil
}
