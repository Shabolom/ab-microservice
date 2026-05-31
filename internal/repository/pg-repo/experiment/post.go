package experiment

import (
	"ab/internal/dto"
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
		return nil, err
	}

	err = s.createLayerExperiments(ctx, tx, created.ID, experiment.LayersID)
	if err != nil {
		return nil, err
	}

	groups, err := s.createExperimentGroups(ctx, tx, created.ID, experiment.Groups)
	if err != nil {
		return nil, err
	}

	created.LayersID = experiment.LayersID
	created.Groups = groups

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
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
			bucket,
			start_date,
			end_date,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING
			id,
			name,
			rollout_percentage,
			bucket,
			start_date,
			end_date,
			status
	`

	var created dto.Experiment

	err := tx.QueryRow(
		ctx,
		query,
		experiment.Name,
		experiment.RolloutPercentage,
		experiment.Bucket,
		experiment.StartDate,
		experiment.EndDate,
		experiment.Status,
	).Scan(
		&created.ID,
		&created.Name,
		&created.RolloutPercentage,
		&created.Bucket,
		&created.StartDate,
		&created.EndDate,
		&created.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("create experiment: %w", err)
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
			return fmt.Errorf("create layer experiment: %w", err)
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
		err := tx.QueryRow(
			ctx,
			query,
			experimentID,
			group.Name,
			group.RollingPercentage,
			group.DeviceID,
		).Scan(&group.ID)
		if err != nil {
			return nil, fmt.Errorf("create experiment group: %w", err)
		}

		createdGroups = append(createdGroups, group)
	}

	return createdGroups, nil
}
