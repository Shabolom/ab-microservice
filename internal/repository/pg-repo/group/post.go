package group

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) Post(ctx context.Context, experimentID int64, group *dto.Group) (*dto.Group, error) {
	query := `
		INSERT INTO experiment_groups (
			experiment_id,
			name,
			rolling_percentage,
			device_ids
		)
		SELECT
			$1,
			$2,
			$3,
			$4
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
		group.DeviceID,
	).Scan(&group.ID)

	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return group, nil
}
