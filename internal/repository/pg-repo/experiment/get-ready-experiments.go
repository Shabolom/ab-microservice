package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetLayersWithExperiments(ctx context.Context) ([]dto.LayerWithExperiments, error) {
	query := `
		WITH active_layers AS (
			SELECT DISTINCT le.layer_id
			FROM experiments e
			JOIN layer_experiments le ON le.experiment_id = e.id
			WHERE e.status = $1
			  AND e.start_date <= NOW()
			  AND e.end_date > NOW()
		)
		SELECT
			le.layer_id,

			e.id,
			e.name,
			e.rollout_percentage,
			le.bucket,
			e.start_date,
			e.end_date,
			e.status
		FROM active_layers al
		JOIN layer_experiments le ON le.layer_id = al.layer_id
		JOIN experiments e ON e.id = le.experiment_id
		ORDER BY le.layer_id, e.id
	`

	rows, err := s.conn.Query(ctx, query, shortcut.ExpStatusReady)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	layerMap := make(map[int64][]dto.LayerExperiment)

	for rows.Next() {
		var layerID int64
		var exp dto.LayerExperiment

		err := rows.Scan(
			&layerID,
			&exp.ID,
			&exp.Name,
			&exp.RolloutPercentage,
			&exp.Buckets,
			&exp.StartDate,
			&exp.EndDate,
			&exp.Status,
		)
		if err != nil {
			return nil, err
		}

		layerMap[layerID] = append(layerMap[layerID], exp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := make([]dto.LayerWithExperiments, 0, len(layerMap))

	for layerID, experiments := range layerMap {
		result = append(result, dto.LayerWithExperiments{
			LayerID:     layerID,
			Experiments: experiments,
		})
	}

	return result, nil
}
