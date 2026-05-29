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
			namespace = COALESCE($3, namespace),
			rolling_percentage = COALESCE($4, rolling_percentage)
		WHERE id = $1
		RETURNING
			id,
			name,
			namespace,
			rolling_percentage
	`

	var experiment dto.Experiment

	err := s.conn.QueryRow(
		ctx,
		query,
		req.ID,
		req.Name,
		req.Namespace,
		req.RollingPercentage,
	).Scan(
		&experiment.Id,
		&experiment.Name,
		&experiment.NameSpase,
		&experiment.RollingPercentage,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("failed to update experiment no rows")
		}

		return nil, err
	}

	return &experiment, nil
}
