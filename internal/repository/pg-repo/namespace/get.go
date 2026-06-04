package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetByID(ctx context.Context, namespaceID int64) (*dto.NameSpace, error) {
	query := `
		SELECT
			id,
			name,
			description
		FROM namespaces
		WHERE id = $1
	`

	var namespace dto.NameSpace

	err := s.conn.QueryRow(
		ctx,
		query,
		namespaceID,
	).Scan(
		&namespace.ID,
		&namespace.Name,
		&namespace.Description,
	)

	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return &namespace, nil
}
