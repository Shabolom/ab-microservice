package featureToggles

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) UpdateRolloutPercentage(ctx context.Context, id int, rolloutPercentage int) error {
	query := `
		UPDATE feature_toggles
		SET
			rollout_percentage = $1,
			updated_at = now()
		WHERE id = $2
	`

	result, err := s.conn.Exec(
		ctx,
		query,
		rolloutPercentage,
		id,
	)

	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if result.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
