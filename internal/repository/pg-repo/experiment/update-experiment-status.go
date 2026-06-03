package experiment

import (
	"context"
	"errors"
)

func (s *Storage) UpdateStatus(ctx context.Context, experimentID int64, status string) error {
	query := `
		UPDATE experiments
		SET status = $2
		WHERE id = $1
	`

	tag, err := s.conn.Exec(
		ctx,
		query,
		experimentID,
		status,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("experiment not found")
	}

	return nil
}
