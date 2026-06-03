package experiment

import (
	"ab/pkg/shortcut"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) UpdateExpired(ctx context.Context, date time.Time) ([]int64, error) {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	query := `
		UPDATE experiments
		SET status = $2
		WHERE end_date < $1
		  AND status = $3
		RETURNING id
	`

	rows, err := tx.Query(
		ctx,
		query,
		date,
		shortcut.ExpStatusEnded,
		shortcut.ExpStatusActive,
	)
	if err != nil {
		return nil, fmt.Errorf("mark expired experiments as ended: %w", err)
	}
	defer rows.Close()

	ids := make([]int64, 0)

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan experiment id: %w", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate experiment ids: %w", err)
	}

	rows.Close()

	for _, id := range ids {
		if err := s.deleteExperimentFromLayerTx(ctx, tx, id); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return ids, nil
}

func (s *Storage) deleteExperimentFromLayerTx(
	ctx context.Context,
	tx pgx.Tx,
	expID int64,
) error {
	query := `
		DELETE FROM layer_experiments
		WHERE experiment_id = $1
	`

	_, err := tx.Exec(ctx, query, expID)
	if err != nil {
		return fmt.Errorf("delete experiment from layer: %w", err)
	}

	return nil
}
