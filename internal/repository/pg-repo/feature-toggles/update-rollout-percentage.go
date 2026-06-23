package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) UpdatePercentage(ctx context.Context, update *dto.FeatureTogglePercentageUpdate) error {
	query := `
		UPDATE feature_toggles
		SET
			rollout_percentage = $1,
			ios = $2,
			android = $3,
			web = $4,
			updated_at = now()
		WHERE id = $5
	`

	result, err := s.conn.Exec(
		ctx,
		query,
		update.Percentage,
		update.Ios,
		update.Android,
		update.Web,
		update.FeatureID,
	)
	if err != nil {
		return shortcut.MapStorageError(err)
	}

	if result.RowsAffected() == 0 {
		return shortcut.ErrNotFound
	}

	return nil
}
