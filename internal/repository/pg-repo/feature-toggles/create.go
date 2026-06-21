package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) Create(ctx context.Context, featureToggle *dto.FeatureToggle) error {
	query := `
		INSERT INTO feature_toggles (
			namespace_id,
			name,
			status
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := s.conn.QueryRow(
		ctx,
		query,
		featureToggle.NamespaceID,
		featureToggle.Name,
		featureToggle.Status,
	).Scan(
		&featureToggle.ID,
		&featureToggle.CreatedAt,
		&featureToggle.UpdatedAt,
	)

	if err != nil {
		return shortcut.MapStorageError(err)
	}

	return nil
}
