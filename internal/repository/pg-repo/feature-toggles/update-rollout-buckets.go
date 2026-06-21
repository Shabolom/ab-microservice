package featureToggles

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) UpdatePercentageAndBuckets(ctx context.Context, id int64, buckets []int64) error {
	rolloutPercentage := len(buckets)

	query := `
		UPDATE feature_toggles
		SET
			buckets = $1,
			rollout_percentage = $2,
			updated_at = now()
		WHERE id = $3
	`

	result, err := s.conn.Exec(
		ctx,
		query,
		buckets,
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
