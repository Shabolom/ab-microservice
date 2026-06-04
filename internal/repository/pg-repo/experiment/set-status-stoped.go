package experiment

import (
	"ab/pkg/shortcut"
	"context"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) SetStatusStopped(ctx context.Context, expID int64) error {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err = s.updateStatusStopped(ctx, tx, expID); err != nil {
		return err
	}

	if err = s.deleteExperimentLayer(ctx, tx, expID); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return shortcut.MapStorageError(err)
	}

	return nil
}

func (s *Storage) updateStatusStopped(
	ctx context.Context,
	tx pgx.Tx,
	expID int64,
) error {
	query := `
		UPDATE experiments
		SET status = $2
		WHERE id = $1
	`

	tag, err := tx.Exec(ctx, query, expID, shortcut.ExpStatusStopped)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if tag.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}

func (s *Storage) deleteExperimentLayer(
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
		return shortcut.MapStorageError(err)
	}

	return nil
}
