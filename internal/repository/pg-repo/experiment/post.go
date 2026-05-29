package experiment

import (
	"ab/internal/dto"
	"context"
)

func (s *Storage) CreateExperiment(ctx context.Context, experiment *dto.Experiment) (*dto.Experiment, error) {
	query := `
		INSERT INTO experiments (
			id,
			name,
			namespace,
			rolling_percentage
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			name,
			namespace,
			rolling_percentage
	`

	var created dto.Experiment

	err := s.conn.QueryRow(
		ctx,
		query,
		experiment.Id,
		experiment.Name,
		experiment.NameSpase,
		experiment.RollingPercentage,
	).Scan(
		&created.Id,
		&created.Name,
		&created.NameSpase,
		&created.RollingPercentage,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}
