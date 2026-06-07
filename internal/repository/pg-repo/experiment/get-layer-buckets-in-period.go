package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetLayerBucketsInPeriod(
	ctx context.Context,
	experiment *dto.Experiment,
) ([]dto.LayerBuckets, error) {
	query := `
		WITH target_layers AS (
			SELECT layer_id
			FROM layer_experiments
			WHERE experiment_id = $1
		)
		SELECT
			tl.layer_id,
			COALESCE(
				array_agg(bucket_value) FILTER (WHERE bucket_value IS NOT NULL),
				'{}'
    		) AS buckets
		FROM target_layers tl
		LEFT JOIN layer_experiments le
			ON le.layer_id = tl.layer_id
		LEFT JOIN experiments e
			ON e.id = le.experiment_id
		   AND e.status IN ($2, $3)
		   AND e.start_date <= $4
		   AND e.end_date >= $5
		   AND e.id <> $1
		LEFT JOIN LATERAL unnest(le.bucket) AS bucket_value
			ON e.id IS NOT NULL
		GROUP BY tl.layer_id
		ORDER BY tl.layer_id
	`

	rows, err := s.conn.Query(
		ctx,
		query,
		experiment.ID,
		shortcut.ExpStatusReady,
		shortcut.ExpStatusActive,
		experiment.EndDate,
		experiment.StartDate,
	)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	result := make([]dto.LayerBuckets, 0)

	for rows.Next() {
		var item dto.LayerBuckets

		err = rows.Scan(
			&item.LayerID,
			&item.Buckets,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		result = append(result, item)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return result, nil
}
