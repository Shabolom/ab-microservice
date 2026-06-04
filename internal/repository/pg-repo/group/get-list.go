package group

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
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
		return nil, shortcut.MapStorageError(err)
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
			return nil, shortcut.MapStorageError(err)
		}

		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return groups, nil
}
