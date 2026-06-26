package featureToggles

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) IsFeatureTogglesEnable(ctx context.Context, id int64) (string, error) {
	query := `
		SELECT status
		FROM feature_toggles
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var status string

	err := s.conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(&status)

	if err != nil {
		return "", shortcut.MapStorageError(err)
	}

	return status, nil
}
