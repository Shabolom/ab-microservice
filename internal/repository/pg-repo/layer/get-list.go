package layer

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetList(ctx context.Context) ([]*dto.Layer, error) {
	query := `
		SELECT
			id,
			namespace_id,
			name,
			description
		FROM layers
		ORDER BY id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	var layers []*dto.Layer

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

		layers = append(layers, &layer)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layers, nil
}
