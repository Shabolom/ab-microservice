package experiment

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) GetReadyToStart(ctx context.Context, startDate time.Time) ([]int64, error) {
	query := `
		SELECT id
		FROM experiments
		WHERE start_date <= $1
		  AND end_date >= $1
	`

	rows, err := s.conn.Query(ctx, query, startDate)
	if err != nil {
		return nil, fmt.Errorf("get ready to start experiments: %w", err)
	}
	defer rows.Close()

	ids := make([]int64, 0)

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan experiment id: %w", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate experiment ids: %w", err)
	}

	return ids, nil
}
