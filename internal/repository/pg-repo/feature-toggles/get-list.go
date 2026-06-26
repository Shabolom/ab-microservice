package featureToggles

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetList(ctx context.Context) ([]*dto.RawFeatureToggle, error) {
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
		WHERE deleted_at IS NULL
		ORDER BY id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	var featureToggles []*dto.RawFeatureToggle

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

		featureToggles = append(featureToggles, &featureToggle)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return featureToggles, nil
}
