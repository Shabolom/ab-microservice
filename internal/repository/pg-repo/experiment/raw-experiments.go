package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
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
			e.passing_cities,
			e.excluded_cities,
			e.passing_stores,
			e.excluded_stores,

			ns.id,
			ns.name,
			ns.description,

			COALESCE(
				(
					SELECT array_agg(DISTINCT b ORDER BY b)
					FROM layer_experiments le2
					CROSS JOIN LATERAL unnest(le2.bucket) AS b
					WHERE le2.experiment_id = e.id
				),
				'{}'::int[]
			) AS bucket,

			COALESCE(
				(
					SELECT jsonb_agg(
						DISTINCT jsonb_build_object(
							'id', l2.id,
							'namespace_id', l2.namespace_id,
							'name', l2.name,
							'description', l2.description
						)
					)
					FROM layer_experiments le2
					JOIN layers l2 ON l2.id = le2.layer_id
					WHERE le2.experiment_id = e.id
				),
				'[]'::jsonb
			) AS layers,

			COALESCE(
				(
					SELECT jsonb_agg(
						jsonb_build_object(
							'id', g.id,
							'name', g.name,
							'rolling_percentage', g.rolling_percentage,
							'device_id', g.device_ids
						)
						ORDER BY g.id
					)
					FROM experiment_groups g
					WHERE g.experiment_id = e.id
				),
				'[]'::jsonb
			) AS groups,

			COALESCE(
				(
					SELECT jsonb_agg(
						jsonb_build_object(
							'id', pg.id,
							'percent', pg.percent,
							'conditions', COALESCE(
								(
									SELECT jsonb_agg(
										jsonb_build_object(
											'id', cpc.id,
											'parameter_id', cpc.parameter_id,
											'parameter_group_id', cpc.parameter_group_id,
											'parameter_type', cp.type,
											'parameter_name', cp.name,
											'value', cpc.value,
											'condition', cpc.condition
										)
										ORDER BY cpc.id
									)
									FROM customparametercondition cpc
									JOIN customparameter cp
										ON cp.id = cpc.parameter_id
									WHERE cpc.parameter_group_id = pg.id
								),
								'[]'::jsonb
							)
						)
						ORDER BY pg.id
					)
					FROM parametergroup pg
					WHERE pg.experiment_id = e.id
				),
				'[]'::jsonb
			) AS custom_params_groups
		FROM experiments e
		JOIN namespaces ns
			ON ns.name = e.namespace
		WHERE ns.name = $1
		ORDER BY e.id
	`

	rows, err := s.conn.Query(ctx, query, namespace)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	result := make([]dto.RawExperiment, 0)

	for rows.Next() {
		var (
			exp dto.RawExperiment

			namespaceID          int64
			namespaceName        string
			namespaceDescription string

			layersRaw             []byte
			groupsRaw             []byte
			customParamsGroupsRaw []byte
		)

		err = rows.Scan(
			&exp.Id,
			&exp.Name,
			&exp.RollingPercentage,
			&exp.Status,
			&exp.StartDate,
			&exp.EndDate,
			&exp.PassingCities,
			&exp.ExcludedCities,
			&exp.PassingStores,
			&exp.ExcludedStores,

			&namespaceID,
			&namespaceName,
			&namespaceDescription,

			&exp.Bucket,

			&layersRaw,
			&groupsRaw,
			&customParamsGroupsRaw,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		exp.NameSpaceName = namespaceName
		exp.NameSpace = dto.NameSpace{
			ID:          namespaceID,
			Name:        namespaceName,
			Description: namespaceDescription,
		}

		if err = json.Unmarshal(layersRaw, &exp.Layer); err != nil {
			return nil, fmt.Errorf("unmarshal layers: %w", err)
		}

		if err = json.Unmarshal(groupsRaw, &exp.Group); err != nil {
			return nil, fmt.Errorf("unmarshal groups: %w", err)
		}

		if err = json.Unmarshal(customParamsGroupsRaw, &exp.CustomParamsGroups); err != nil {
			return nil, fmt.Errorf("unmarshal custom params groups: %w", err)
		}

		result = append(result, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return result, nil
}
