package experiment

import (
	"context"
	"fmt"
	"time"
)

func (s *Storage) GetExpired(ctx context.Context, date time.Time) ([]int64, error) {
	query := `
		SELECT id
		FROM experiments
		WHERE end_date < $1
	`

	rows, err := s.conn.Query(ctx, query, date)
	if err != nil {
		return nil, fmt.Errorf("get expired experiments: %w", err)
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
