package experement

import (
	"ab/internal/dto"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetExperiment(ctx context.Context, id string) (*dto.Experiment, error) {
	query := `
		SELECT
			id,
			name,
			namespace,
			rolling_percentage
		FROM experiments
		WHERE id = $1
	`

	var experiment dto.Experiment

	err := s.conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&experiment.Id,
		&experiment.Name,
		&experiment.NameSpase,
		&experiment.RollingPercentage,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("experiment not found")
		}

		return nil, err
	}

	return &experiment, nil
}
