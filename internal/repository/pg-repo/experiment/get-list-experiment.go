package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"encoding/json"
	"fmt"
)

func (s *Storage) GetList(ctx context.Context) ([]dto.Experiment, error) {
	query := `
		SELECT
			e.id,
			e.name,
			e.namespace,
			e.rollout_percentage,
			e.status,
			e.start_date,
			e.end_date,
			e.passing_cities,
			e.excluded_cities,
			e.passing_stores,
			e.excluded_stores,

			COALESCE(
				(
					SELECT array_agg(le.layer_id ORDER BY le.layer_id)
					FROM layer_experiments le
					WHERE le.experiment_id = e.id
				),
				'{}'::bigint[]
			) AS layers_id,

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
											'value', cpc.value,
											'condition', cpc.condition
										)
										ORDER BY cpc.id
									)
									FROM customparametercondition cpc
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
			) AS params_groups,

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
		ORDER BY e.id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	experiments := make([]dto.Experiment, 0)

	for rows.Next() {
		var (
			experiment      dto.Experiment
			paramsGroupsRaw []byte
			groupsRaw       []byte
		)

		err = rows.Scan(
			&experiment.ID,
			&experiment.Name,
			&experiment.Namespace,
			&experiment.RolloutPercentage,
			&experiment.Status,
			&experiment.StartDate,
			&experiment.EndDate,
			&experiment.PassingCities,
			&experiment.ExcludedCities,
			&experiment.PassingStores,
			&experiment.ExcludedStores,
			&experiment.LayersID,
			&paramsGroupsRaw,
			&groupsRaw,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		if err = json.Unmarshal(paramsGroupsRaw, &experiment.ParamsGroups); err != nil {
			return nil, fmt.Errorf("unmarshal params groups: %w", err)
		}

		if err = json.Unmarshal(groupsRaw, &experiment.Groups); err != nil {
			return nil, fmt.Errorf("unmarshal groups: %w", err)
		}

		experiments = append(experiments, experiment)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return experiments, nil
}
