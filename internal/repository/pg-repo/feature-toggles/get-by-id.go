package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetByID(ctx context.Context, id int64) (*dto.RawFeatureToggle, error) {
	query := `
		SELECT
			id,
			rollout_percentage,
			namespace_id,
			name,
			buckets,
			status,
			created_at,
			updated_at,
			deleted_at
		FROM feature_toggles
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	featureToggle := &dto.RawFeatureToggle{}

	err := s.conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&featureToggle.ID,
		&featureToggle.RolloutPercentage,
		&featureToggle.NamespaceID,
		&featureToggle.Name,
		&featureToggle.Buckets,
		&featureToggle.Status,
		&featureToggle.CreatedAt,
		&featureToggle.UpdatedAt,
		&featureToggle.DeletedAt,
	)

	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return featureToggle, nil
}
