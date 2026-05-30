package experiment

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error) {
	query := `
		SELECT
			e.id,
			e.name,
			e.rollout_percentage,
			e.start_date,
			e.end_date,

			ns.id,
			ns.name,
			ns.description,

			l.id,
			l.namespace_id,
			l.name,
			l.description,

			g.id,
			g.name,
			g.rolling_percentage
		FROM experiments e
		LEFT JOIN layer_experiments le
			ON le.experiment_id = e.id
		LEFT JOIN layers l
			ON l.id = le.layer_id
		LEFT JOIN namespaces ns
			ON ns.id = l.namespace_id
		LEFT JOIN experiment_groups g
			ON g.experiment_id = e.id
		WHERE ns.name = $1
		ORDER BY e.id, l.id, g.id
	`

	rows, err := s.conn.Query(ctx, query, namespace)
	if err != nil {
		return nil, fmt.Errorf("get raw experiments: %w", err)
	}
	defer rows.Close()

	result := make([]dto.RawExperiment, 0)

	for rows.Next() {
		var (
			exp dto.RawExperiment

			namespaceID          *int64
			namespaceName        *string
			namespaceDescription *string

			layerID          *int64
			layerNamespaceID *int64
			layerName        *string
			layerDescription *string

			groupID                *int64
			groupName              *string
			groupRollingPercentage *int
		)

		err = rows.Scan(
			&exp.Id,
			&exp.Name,
			&exp.RollingPercentage,
			&exp.StartDate,
			&exp.EndDate,

			&namespaceID,
			&namespaceName,
			&namespaceDescription,

			&layerID,
			&layerNamespaceID,
			&layerName,
			&layerDescription,

			&groupID,
			&groupName,
			&groupRollingPercentage,
		)
		if err != nil {
			return nil, fmt.Errorf("scan raw experiment: %w", err)
		}

		if namespaceID != nil {
			exp.NameSpaceName = *namespaceName
			exp.NameSpace = dto.NameSpace{
				ID:          *namespaceID,
				Name:        *namespaceName,
				Description: *namespaceDescription,
			}
		}

		if layerID != nil {
			exp.Layer = append(exp.Layer, dto.Layer{
				ID:          *layerID,
				NameSpaceID: *layerNamespaceID,
				Name:        *layerName,
				Description: *layerDescription,
			})
		}

		if groupID != nil {
			exp.Group = append(exp.Group, dto.Group{
				ID:                *groupID,
				Name:              *groupName,
				RollingPercentage: *groupRollingPercentage,
			})
		}

		result = append(result, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate raw experiments: %w", err)
	}

	return result, nil
}
