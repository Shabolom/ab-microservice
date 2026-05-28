package experement

import (
	"ab/internal/dto"
	"context"
)

func (s *Storage) GetList(ctx context.Context) ([]*dto.Experiment, error) {
	query := `
		SELECT
			id,
			name,
			namespace,
			rolling_percentage
		FROM experiments
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	experiments := make([]*dto.Experiment, 0)

	for rows.Next() {
		experiment := &dto.Experiment{}

		err = rows.Scan(
			&experiment.Id,
			&experiment.Name,
			&experiment.NameSpase,
			&experiment.RollingPercentage,
		)
		if err != nil {
			return nil, err
		}

		experiments = append(experiments, experiment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return experiments, nil
}
