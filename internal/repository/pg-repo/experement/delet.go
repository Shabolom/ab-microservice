package experement

import (
	"context"
	"errors"
)

func (s *Storage) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM experiments
		WHERE id = $1
	`

	tag, err := s.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("experiment does not exist")
	}

	return nil
}
