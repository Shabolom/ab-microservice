package layer

import (
	"ab/pkg/shortcut"
	"context"
	"fmt"
)

func (s *Storage) GetIDsByNamespaceId(ctx context.Context, namespaceId int64) ([]int64, error) {
	query := `
		SELECT id
		FROM layers
		WHERE namespace_id = $1
		ORDER BY id
	`

	rows, err := s.conn.Query(ctx, query, namespaceId)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	layerIDs := make([]int64, 0)

	for rows.Next() {
		var layerID int64

		if err := rows.Scan(&layerID); err != nil {
			return nil, fmt.Errorf("scan layer id: %w", err)
		}

		layerIDs = append(layerIDs, layerID)
	}

	if err := rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layerIDs, nil
}
