package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) SetStatusDisabled(ctx context.Context, id int64) error {
	query := `
		UPDATE feature_toggles
		SET
			status = $1,
			buckets = '{}',
			updated_at = now()
		WHERE id = $2
	`

	result, err := s.conn.Exec(ctx, query, dto.FeatureToggleStatusDisabled, id)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if result.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
