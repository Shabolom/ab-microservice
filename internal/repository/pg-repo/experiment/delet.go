package experiment

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM experiments
		WHERE id = $1
	`

	tag, err := s.conn.Exec(ctx, query, id)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if tag.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
