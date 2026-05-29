package experiment

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) GetRawExperiments(ctx context.Context) ([]*dto.RawExperiment, error) {
	query := `
		SELECT
			e.id,
			e.name,
			e.namespace,
			e.rolling_percentage,

			g.id,
			g.name,
			g.rolling_percentage
		FROM experiments e
		LEFT JOIN experiment_groups g
			ON g.experiment_id = e.id
		ORDER BY e.id, g.id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get raw experiments: %w", err)
	}
	defer rows.Close()

	experimentsMap := make(map[int64]*dto.RawExperiment)
	experimentIDs := make([]int64, 0)

	for rows.Next() {
		var (
			expID                int64
			expName              string
			expNamespace         string
			expRollingPercentage int

			groupID                *int64
			groupName              *string
			groupRollingPercentage *int
		)

		err = rows.Scan(
			&expID,
			&expName,
			&expNamespace,
			&expRollingPercentage,
			&groupID,
			&groupName,
			&groupRollingPercentage,
		)
		if err != nil {
			return nil, fmt.Errorf("scan raw experiment: %w", err)
		}

		exp, exists := experimentsMap[expID]
		if !exists {
			exp = &dto.RawExperiment{
				Id:                expID,
				Name:              expName,
				NameSpase:         expNamespace,
				RollingPercentage: expRollingPercentage,
				Group:             make([]dto.Group, 0),
			}

			experimentsMap[expID] = exp
			experimentIDs = append(experimentIDs, expID)
		}

		if groupID != nil {
			exp.Group = append(exp.Group, dto.Group{
				ID:                *groupID,
				Name:              *groupName,
				RollingPercentage: *groupRollingPercentage,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate raw experiments: %w", err)
	}

	result := make([]*dto.RawExperiment, 0, len(experimentsMap))
	for _, id := range experimentIDs {
		result = append(result, experimentsMap[id])
	}

	return result, nil
}
