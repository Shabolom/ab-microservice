package group

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) Update(ctx context.Context, group *dto.UpdateGroup) (*dto.Group, error) {
	query := `
		UPDATE experiment_groups
		SET
			name = COALESCE($1, name),
			rolling_percentage = COALESCE($2, rolling_percentage)
		WHERE id = $3
		RETURNING
			id,
			name,
			rolling_percentage
	`

	updatedGroup := &dto.Group{}

	err := s.conn.QueryRow(
		ctx,
		query,
		group.Name,
		group.RollingPercentage,
		group.ID,
	).Scan(
		&updatedGroup.ID,
		&updatedGroup.Name,
		&updatedGroup.RollingPercentage,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shortcut.MapStorageError(err)
		}

		return nil, shortcut.MapStorageError(err)
	}

	return updatedGroup, nil
}
