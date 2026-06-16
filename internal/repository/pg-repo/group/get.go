package group

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetByID(ctx context.Context, id int64) (*dto.Group, error) {
	query := `
		SELECT
			id,
			name,
			rolling_percentage
		FROM experiment_groups
		WHERE id = $1
	`

	group := &dto.Group{}

	err := s.conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&group.ID,
		&group.Name,
		&group.RollingPercentage,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shortcut.MapStorageError(err)
		}

		return nil, shortcut.MapStorageError(err)
	}

	return group, nil
}
