package experiment

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (s *Storage) GetExperiment(ctx context.Context, id int64) (*dto.Experiment, error) {
	query := `
		SELECT
			id,
			name,
			rollout_percentage,
			start_date,
			end_date,
			status,
			namespace
		FROM experiments
		WHERE id = $1
	`

	var experiment dto.Experiment

	err := s.conn.QueryRow(ctx, query, id).Scan(
		&experiment.ID,
		&experiment.Name,
		&experiment.RolloutPercentage,
		&experiment.StartDate,
		&experiment.EndDate,
		&experiment.Status,
		&experiment.NameSpace,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, shortcut.MapStorageError(err)
		}

		return nil, err
	}

	return &experiment, nil
}
