package layer

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetByID(ctx context.Context, id int64) (*dto.Layer, error) {
	query := `
		SELECT
			id,
			namespace_id,
			name,
			description
		FROM layers
		WHERE id = $1
	`

	layer := &dto.Layer{}

	err := s.conn.QueryRow(ctx, query, id).Scan(
		&layer.ID,
		&layer.NameSpaceID,
		&layer.Name,
		&layer.Description,
	)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return layer, nil
}
