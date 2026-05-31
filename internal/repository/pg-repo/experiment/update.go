package experiment

import (
	"ab/internal/dto"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) Update(ctx context.Context, req *dto.UpdateExperiment) (*dto.Experiment, error) {
	query := `
		UPDATE experiments
		SET
			name = COALESCE($2, name),
			rollout_percentage = COALESCE($3, rollout_percentage),
			bucket = COALESCE($4, bucket),
			start_date = COALESCE($5, start_date),
			end_date = COALESCE($6, end_date),
			status = COALESCE($7, status)
		WHERE id = $1
		RETURNING
			id,
			name,
			rollout_percentage,
			bucket,
			start_date,
			end_date,
			status
	`

	var experiment dto.Experiment

	err := s.conn.QueryRow(
		ctx,
		query,
		req.ID,
		req.Name,
		req.RolloutPercentage,
		req.Bucket,
		req.StartDate,
		req.EndDate,
		req.Status,
	).Scan(
		&experiment.ID,
		&experiment.Name,
		&experiment.RolloutPercentage,
		&experiment.Bucket,
		&experiment.StartDate,
		&experiment.EndDate,
		&experiment.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("experiment not found")
		}

		return nil, err
	}

	return &experiment, nil
}
