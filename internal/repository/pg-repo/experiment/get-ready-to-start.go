package experiment

import (
	"ab/pkg/shortcut"
	"context"
	"fmt"
	"time"
)

func (s *Storage) UpdateReadyToStart(ctx context.Context, startDate time.Time) ([]int64, error) {
	query := `
		UPDATE experiments
		SET status = $2
		WHERE start_date <= $1
		  AND end_date >= $1
		  AND status = $3
		RETURNING id
	`

	rows, err := s.conn.Query(
		ctx,
		query,
		startDate,
		shortcut.ExpStatusActive,
		shortcut.ExpStatusReady,
	)
	if err != nil {
		return nil, fmt.Errorf("activate ready experiments: %w", err)
	}
	defer rows.Close()

	ids := make([]int64, 0)

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan activated experiment id: %w", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate activated experiment ids: %w", err)
	}

	return ids, nil
}
