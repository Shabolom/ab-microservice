package customParams

import (
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) CountDistinctNamespaces(ctx context.Context, ids []int64) (int64, error) {
	query := `
		SELECT COUNT(DISTINCT namespace_id)
		FROM customparameter
		WHERE id = ANY($1)
	`

	var count int64

	err := s.conn.QueryRow(ctx, query, ids).Scan(&count)
	if err != nil {
		return 0, shortcut.MapStorageError(err)
	}

	return count, nil
}
