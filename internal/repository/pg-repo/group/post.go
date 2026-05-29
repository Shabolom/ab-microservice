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
		SELECT
			$1,
			$2,
			$3
		WHERE (
			SELECT COALESCE(SUM(rolling_percentage), 0)
			FROM experiment_groups
			WHERE experiment_id = $1
		) + $3 <= 100
		RETURNING id;
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
