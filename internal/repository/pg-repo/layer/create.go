package layer

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) Create(ctx context.Context, layer *dto.Layer) error {
	query := `
		INSERT INTO layers (
			namespace_id,
			name,
			description
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := s.conn.QueryRow(
		ctx,
		query,
		layer.NameSpaceID,
		layer.Name,
		layer.Description,
	).Scan(&layer.ID)
	if err != nil {
		return fmt.Errorf("create layer: %w", err)
	}

	return nil
}
