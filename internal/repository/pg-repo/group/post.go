package group

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) Post(ctx context.Context, experimentID string, group *dto.Group) (*dto.Group, error) {
	query := `
		INSERT INTO experiment_groups (
			experiment_id,
			name,
			rolling_percentage
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := s.conn.QueryRow(
		ctx,
		query,
		experimentID,
		group.Name,
		group.RollingPercentage,
	).Scan(&group.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return group, nil
}
