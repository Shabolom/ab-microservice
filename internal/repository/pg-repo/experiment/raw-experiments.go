package experiment

import (
	"ab/internal/dto"
	"context"
	"encoding/json"
	"fmt"
)

func (s *Storage) GetRawExperiments(ctx context.Context, namespace string) ([]dto.RawExperiment, error) {
	query := `
		SELECT
			e.id,
			e.name,
			e.rollout_percentage,
			e.status,
			e.start_date,
			e.end_date,

			ns.id,
			ns.name,
			ns.description,

			COALESCE(
				jsonb_agg(
					DISTINCT jsonb_build_object(
						'id', l.id,
						'namespace_id', l.namespace_id,
						'name', l.name,
						'description', l.description
					)
				) FILTER (WHERE l.id IS NOT NULL),
				'[]'::jsonb
			) AS layers,

			COALESCE(
				jsonb_agg(
					DISTINCT jsonb_build_object(
						'id', g.id,
						'name', g.name,
						'rolling_percentage', g.rolling_percentage,
						'device_id', g.device_ids
					)
				) FILTER (WHERE g.id IS NOT NULL),
				'[]'::jsonb
			) AS groups
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
		GROUP BY
			e.id,
			e.name,
			e.rollout_percentage,
			e.status,
			e.start_date,
			e.end_date,
			ns.id,
			ns.name,
			ns.description
		ORDER BY e.id
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

			namespaceID          int64
			namespaceName        string
			namespaceDescription string

			layersRaw []byte
			groupsRaw []byte
		)

		err = rows.Scan(
			&exp.Id,
			&exp.Name,
			&exp.RollingPercentage,
			&exp.Status,
			&exp.StartDate,
			&exp.EndDate,

			&namespaceID,
			&namespaceName,
			&namespaceDescription,

			&layersRaw,
			&groupsRaw,
		)
		if err != nil {
			return nil, fmt.Errorf("scan raw experiment: %w", err)
		}

		exp.NameSpaceName = namespaceName
		exp.NameSpace = dto.NameSpace{
			ID:          namespaceID,
			Name:        namespaceName,
			Description: namespaceDescription,
		}

		if err = json.Unmarshal(layersRaw, &exp.Layer); err != nil {
			return nil, fmt.Errorf("unmarshal layers for experiment %d: %w", exp.Id, err)
		}

		if err = json.Unmarshal(groupsRaw, &exp.Group); err != nil {
			return nil, fmt.Errorf("unmarshal groups for experiment %d: %w", exp.Id, err)
		}

		result = append(result, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate raw experiments: %w", err)
	}

	return result, nil
}
