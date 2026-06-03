package experiment

import (
	"ab/internal/dto"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) UpdateStatusBucketsTx(
	ctx context.Context,
	experimentID int64,
	status string,
	layerBuckets []dto.LayerBuckets,
) error {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
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
			return err
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
		UPDATE layer_experiments
		SET bucket = $3
		WHERE layer_id = $1
		  AND experiment_id = $2
	`

	_, err := tx.Exec(ctx, query, layerID, experimentID, buckets)
	if err != nil {
		return fmt.Errorf("update layer buckets: %w", err)
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
		return fmt.Errorf("update experiment status: %w", err)
	}

	return nil
}
