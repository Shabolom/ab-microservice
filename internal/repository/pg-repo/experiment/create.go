package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) CreateExperiment(ctx context.Context, experiment *dto.Experiment) (*dto.Experiment, error) {
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	created, err := s.createExperiment(ctx, tx, experiment)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	err = s.createLayerExperiments(ctx, tx, created.ID, experiment.LayersID)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	groups, err := s.createExperimentGroups(ctx, tx, created.ID, experiment.Groups)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	paramsGroups, err := s.createParamGroups(ctx, tx, created.ID, experiment.ParamsGroups)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	created.LayersID = experiment.LayersID
	created.Groups = groups
	created.ParamsGroups = paramsGroups

	if err = tx.Commit(ctx); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return created, nil
}

func (s *Storage) createExperiment(
	ctx context.Context,
	tx pgx.Tx,
	experiment *dto.Experiment,
) (*dto.Experiment, error) {
	query := `
		INSERT INTO experiments (
			name,
			rollout_percentage,
			start_date,
			end_date,
			status,
			namespace,
			passing_cities,
			excluded_cities,
			passing_stores,
			excluded_stores
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING
			id,
			name,
			rollout_percentage,
			start_date,
			end_date,
			status,
			namespace,
			passing_cities,
			excluded_cities,
			passing_stores,
			excluded_stores
	`

	var created dto.Experiment

	err := tx.QueryRow(
		ctx,
		query,
		experiment.Name,
		experiment.RolloutPercentage,
		experiment.StartDate,
		experiment.EndDate,
		experiment.Status,
		experiment.NameSpace,
		experiment.PassingCities,
		experiment.ExcludedCities,
		experiment.PassingStores,
		experiment.ExcludedStores,
	).Scan(
		&created.ID,
		&created.Name,
		&created.RolloutPercentage,
		&created.StartDate,
		&created.EndDate,
		&created.Status,
		&created.NameSpace,
		&created.PassingCities,
		&created.ExcludedCities,
		&created.PassingStores,
		&created.ExcludedStores,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *Storage) createLayerExperiments(
	ctx context.Context,
	tx pgx.Tx,
	experimentID int64,
	layerIDs []int64,
) error {
	query := `
		INSERT INTO layer_experiments (
			layer_id,
			experiment_id
		)
		VALUES ($1, $2)
	`

	for _, layerID := range layerIDs {
		_, err := tx.Exec(ctx, query, layerID, experimentID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) createExperimentGroups(
	ctx context.Context,
	tx pgx.Tx,
	experimentID int64,
	groups []dto.Group,
) ([]dto.Group, error) {
	query := `
		INSERT INTO experiment_groups (
			experiment_id,
			name,
			rolling_percentage,
			device_ids
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	createdGroups := make([]dto.Group, 0, len(groups))

	for _, group := range groups {
		if group.DeviceID == nil {
			group.DeviceID = []int64{}
		}

		err := tx.QueryRow(
			ctx,
			query,
			experimentID,
			group.Name,
			group.RollingPercentage,
			group.DeviceID,
		).Scan(&group.ID)
		if err != nil {
			return nil, err
		}

		createdGroups = append(createdGroups, group)
	}

	return createdGroups, nil
}

func (s *Storage) createParamGroups(
	ctx context.Context,
	tx pgx.Tx,
	experimentID int64,
	paramGroups []dto.ParamGroup,
) ([]dto.ParamGroup, error) {
	query := `
		INSERT INTO parametergroup (
			percent,
			experiment_id
		)
		VALUES ($1, $2)
		RETURNING id
	`

	createdGroups := make([]dto.ParamGroup, 0, len(paramGroups))

	for _, group := range paramGroups {
		err := tx.QueryRow(
			ctx,
			query,
			group.Percent,
			experimentID,
		).Scan(&group.ID)
		if err != nil {
			return nil, err
		}

		conditions, err := s.createParamGroupConditions(
			ctx,
			tx,
			group.ID,
			group.ParamsWithConditions,
		)
		if err != nil {
			return nil, err
		}

		group.ParamsWithConditions = conditions
		createdGroups = append(createdGroups, group)
	}

	return createdGroups, nil
}

func (s *Storage) createParamGroupConditions(
	ctx context.Context,
	tx pgx.Tx,
	paramGroupID int64,
	conditions []dto.CustomParamWithCondition,
) ([]dto.CustomParamWithCondition, error) {
	query := `
		INSERT INTO customparametercondition (
			parameter_id,
			value,
			condition,
			parameter_group_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			parameter_id,
			value,
			condition,
			parameter_group_id
	`

	createdConditions := make([]dto.CustomParamWithCondition, 0, len(conditions))

	for _, condition := range conditions {
		var created dto.CustomParamWithCondition

		err := tx.QueryRow(
			ctx,
			query,
			condition.ParameterID,
			condition.Value,
			condition.Condition,
			paramGroupID,
		).Scan(
			&created.ID,
			&created.ParameterID,
			&created.Value,
			&created.Condition,
			&created.ParameterGroupID,
		)
		if err != nil {
			return nil, err
		}

		createdConditions = append(createdConditions, created)
	}

	return createdConditions, nil
}
