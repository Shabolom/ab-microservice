package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) Create(ctx context.Context, namespace *dto.NameSpace) (*dto.NameSpace, error) {
	query := `
		INSERT INTO namespaces (
			name,
			description
		)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int64

	err := s.conn.QueryRow(
		ctx,
		query,
		namespace.Name,
		namespace.Description,
	).Scan(&id)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	namespace.ID = id

	return namespace, nil
}
