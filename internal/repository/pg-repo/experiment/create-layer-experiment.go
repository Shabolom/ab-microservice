package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) CreateAndGetLayerExperiment(ctx context.Context, layerIDs []int64, expID int64) ([]dto.LayerBuckets, error) {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = s.createLayerExperimentTx(ctx, tx, layerIDs, expID)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	layerBuckets, err := s.getLayerBucketsByExperimentTx(ctx, tx, expID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layerBuckets, nil
}

func (s *Storage) createLayerExperimentTx(
	ctx context.Context,
	tx pgx.Tx,
	layerIDs []int64,
	expID int64,
) error {
	query := `
		INSERT INTO layer_experiments (
			layer_id,
			experiment_id
		)
		SELECT
			unnest($1::bigint[]),
			$2
		ON CONFLICT (layer_id, experiment_id)
		DO NOTHING
	`

	_, err := tx.Exec(ctx, query, layerIDs, expID)
	if err != nil {
		return fmt.Errorf("create layer experiment: %w", err)
	}

	return nil
}

func (s *Storage) getLayerBucketsByExperimentTx(
	ctx context.Context,
	tx pgx.Tx,
	expID int64,
) ([]dto.LayerBuckets, error) {
	query := `
		SELECT
			layer_id,
			bucket
		FROM layer_experiments
		WHERE experiment_id = $1
		ORDER BY layer_id
	`

	rows, err := tx.Query(ctx, query, expID)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	layerBuckets := make([]dto.LayerBuckets, 0)

	for rows.Next() {
		var item dto.LayerBuckets

		err = rows.Scan(
			&item.LayerID,
			&item.Buckets,
		)
		if err != nil {
			return nil, fmt.Errorf("scan layer buckets: %w", err)
		}

		layerBuckets = append(layerBuckets, item)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layerBuckets, nil
}
