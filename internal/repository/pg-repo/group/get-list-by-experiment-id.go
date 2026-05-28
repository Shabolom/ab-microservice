package group

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) GetListByExperimentID(ctx context.Context, experimentID string) ([]*dto.Group, error) {
	query := `
		SELECT
			id,
			name,
			rolling_percentage
		FROM experiment_groups
		WHERE experiment_id = $1
		ORDER BY id
	`

	rows, err := s.conn.Query(
		ctx,
		query,
		experimentID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get groups list: %w", err)
	}
	defer rows.Close()

	var groups []*dto.Group

	for rows.Next() {
		group := &dto.Group{}

		err = rows.Scan(
			&group.ID,
			&group.Name,
			&group.RollingPercentage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group: %w", err)
		}

		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return groups, nil
}
