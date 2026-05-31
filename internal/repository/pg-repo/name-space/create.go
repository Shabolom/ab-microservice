package nameSpace

import (
	"ab/internal/dto"
	"context"
	"fmt"
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
		return nil, fmt.Errorf("create namespace: %w", err)
	}

	namespace.ID = id

	return namespace, nil
}
