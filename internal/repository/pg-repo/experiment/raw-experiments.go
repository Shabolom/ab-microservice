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
			e.namespace,
			e.rollout_percentage,
			e.status,
			e.start_date,
			e.end_date,

			ns.id,
			ns.name,
			ns.description,

			e.passing_cities,
			e.excluded_cities,
			e.passing_stores,
			e.excluded_stores,

			COALESCE(
				(
					SELECT array_agg(DISTINCT b::bigint ORDER BY b::bigint)
					FROM layer_experiments le
					CROSS JOIN LATERAL unnest(le.bucket) AS b
					WHERE le.experiment_id = e.id
				),
				'{}'::bigint[]
			) AS bucket,

			COALESCE(
				(
					SELECT jsonb_agg(
						jsonb_build_object(
							'id', l.id,
							'namespace_id', l.namespace_id,
							'name', l.name,
							'description', l.description
						)
						ORDER BY l.id
					)
					FROM layer_experiments le
					JOIN layers l ON l.id = le.layer_id
					WHERE le.experiment_id = e.id
				),
				'[]'::jsonb
			) AS layers,

			COALESCE(
				(
					SELECT jsonb_agg(
						jsonb_build_object(
							'id', pg.id,
							'percent', pg.percent,
							'params_with_conditions', COALESCE(
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
			) AS custom_params_groups,

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
			) AS groups

		FROM experiments e
		JOIN namespaces ns
			ON ns.name = e.namespace
		WHERE e.namespace = $1
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
			customParamsGroupsRaw []byte
			groupsRaw             []byte
		)

		err = rows.Scan(
			&exp.Id,
			&exp.Name,
			&exp.NamespaceName,
			&exp.RollingPercentage,
			&exp.Status,
			&exp.StartDate,
			&exp.EndDate,

			&namespaceID,
			&namespaceName,
			&namespaceDescription,

			&exp.PassingCities,
			&exp.ExcludedCities,
			&exp.PassingStores,
			&exp.ExcludedStores,

			&exp.Bucket,

			&layersRaw,
			&customParamsGroupsRaw,
			&groupsRaw,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		exp.NameSpace = dto.NameSpace{
			ID:          namespaceID,
			Name:        namespaceName,
			Description: namespaceDescription,
		}

		if err = json.Unmarshal(layersRaw, &exp.Layer); err != nil {
			return nil, fmt.Errorf("unmarshal layers: %w", err)
		}

		if err = json.Unmarshal(customParamsGroupsRaw, &exp.CustomParamsGroups); err != nil {
			return nil, fmt.Errorf("unmarshal custom params groups: %w", err)
		}

		if err = json.Unmarshal(groupsRaw, &exp.Group); err != nil {
			return nil, fmt.Errorf("unmarshal groups: %w", err)
		}

		result = append(result, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return result, nil
}
