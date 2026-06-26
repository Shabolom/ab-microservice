package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetActiveByNamespaceID(ctx context.Context, id int64) ([]dto.RawFeatureToggle, error) {
	query := `
		SELECT
			id,
			namespace_id,
			name,
			status,
			rollout_percentage,
			ios,
			android,
			web,
			created_at,
			updated_at,
			deleted_at
		FROM feature_toggles
		WHERE namespace_id = $1
		  AND status = 'active'
		  AND deleted_at IS NULL
	`

	rows, err := s.conn.Query(ctx, query, id)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	featureToggles := make([]dto.RawFeatureToggle, 0)

	for rows.Next() {
		var featureToggle dto.RawFeatureToggle

		err = rows.Scan(
			&featureToggle.ID,
			&featureToggle.NamespaceID,
			&featureToggle.Name,
			&featureToggle.Status,
			&featureToggle.RolloutPercentage,
			&featureToggle.IOS,
			&featureToggle.Android,
			&featureToggle.Web,
			&featureToggle.CreatedAt,
			&featureToggle.UpdatedAt,
			&featureToggle.DeletedAt,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		featureToggles = append(featureToggles, featureToggle)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return featureToggles, nil
}
