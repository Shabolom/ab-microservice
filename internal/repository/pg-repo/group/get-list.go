package group

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"
)

func (s *Storage) GetList(ctx context.Context, limit int, id int) ([]*dto.Group, error) {
	query := `
		SELECT
			id,
			name,
			rolling_percentage
		FROM experiment_groups
		WHERE id > $1
		ORDER BY id
		LIMIT $2
	`

	rows, err := s.conn.Query(
		ctx,
		query,
		id,
		limit,
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
			return nil, fmt.Errorf("%w: %v", shortcut.ErrFailedToScanGroup, err)
		}

		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", shortcut.ErrRowsIteration, err)
	}

	return groups, nil
}
