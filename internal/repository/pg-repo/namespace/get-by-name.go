package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetByName(ctx context.Context, name string) (*dto.NameSpace, error) {
	query := `
		SELECT
			id,
			name,
			description
		FROM namespaces
		WHERE name = $1
	`

	var namespace dto.NameSpace

	err := s.conn.QueryRow(ctx, query, name).Scan(
		&namespace.ID,
		&namespace.Name,
		&namespace.Description,
	)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return &namespace, nil
}
