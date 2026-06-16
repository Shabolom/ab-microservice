package layer

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetLayers(ctx context.Context, layerIDs []int64) ([]dto.Layer, error) {
	query := `
		SELECT
			id,
			namespace_id,
			name,
			description
		FROM layers
		WHERE id = ANY($1)
	`

	rows, err := s.conn.Query(ctx, query, layerIDs)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	layers := make([]dto.Layer, 0, len(layerIDs))

	for rows.Next() {
		var layer dto.Layer

		err = rows.Scan(
			&layer.ID,
			&layer.NameSpaceID,
			&layer.Name,
			&layer.Description,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		layers = append(layers, layer)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layers, nil
}
