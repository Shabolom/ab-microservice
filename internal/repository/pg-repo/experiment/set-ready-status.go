package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) SetReadyStatus(
	ctx context.Context,
	experimentID int64,
	status string,
	layerBuckets []dto.LayerBuckets,
) error {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	for _, item := range layerBuckets {
		err = s.updateLayerBucketsTx(
			ctx,
			tx,
			item.LayerID,
			experimentID,
			item.Buckets,
		)
		if err != nil {
			return shortcut.MapStorageError(err)
		}
	}

	err = s.updateStatusTx(ctx, tx, experimentID, status)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *Storage) updateLayerBucketsTx(
	ctx context.Context,
	tx pgx.Tx,
	layerID int64,
	experimentID int64,
	buckets []int64,
) error {
	query := `
		INSERT INTO layer_experiments (
			layer_id,
			experiment_id,
			bucket
		)
		VALUES ($1, $2, $3)
		ON CONFLICT (layer_id, experiment_id)
		DO UPDATE SET
			bucket = EXCLUDED.bucket
	`

	_, err := tx.Exec(
		ctx,
		query,
		layerID,
		experimentID,
		buckets,
	)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	return nil
}

func (s *Storage) updateStatusTx(
	ctx context.Context,
	tx pgx.Tx,
	experimentID int64,
	status string,
) error {
	query := `
		UPDATE experiments
		SET status = $2
		WHERE id = $1
	`

	_, err := tx.Exec(ctx, query, experimentID, status)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	return nil
}
