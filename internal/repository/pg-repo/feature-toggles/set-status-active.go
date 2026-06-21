package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) SetStatusActive(ctx context.Context, id int64, buckets []int64) error {
	query := `
		UPDATE feature_toggles
		SET
			status = $1,
			buckets = $2,
			updated_at = now()
		WHERE id = $3
	`

	result, err := s.conn.Exec(
		ctx,
		query,
		dto.FeatureToggleStatusActive,
		buckets,
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
