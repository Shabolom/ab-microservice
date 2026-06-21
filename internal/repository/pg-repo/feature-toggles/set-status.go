package featureToggles

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) SetStatus(ctx context.Context, id int, status string) error {
	query := `
		UPDATE feature_toggles
		SET
			status = $1,
			updated_at = now()
		WHERE id = $2
		RETURNING updated_at
	`

	result, err := s.conn.Exec(
		ctx,
		query,
		status,
		id,
	)

	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if result.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
