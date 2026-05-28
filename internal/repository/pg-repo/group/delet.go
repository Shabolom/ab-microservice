package group

import (
	"ab/pkg/shortcut"
	"context"
	"fmt"
)

func (s *Storage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM experiment_groups
		WHERE id = $1
	`

	tag, err := s.conn.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", shortcut.ErrFailedToDeleteGroup, err)
	}

	if tag.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
