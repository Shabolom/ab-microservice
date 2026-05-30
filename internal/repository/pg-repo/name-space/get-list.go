package nameSpace

import (
	"ab/internal/dto"
	"context"
	"fmt"
)

func (s *Storage) GetList(ctx context.Context) ([]dto.NameSpace, error) {
	query := `
		SELECT
			id,
			name,
			description
		FROM namespaces
		ORDER BY id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get namespaces: %w", err)
	}
	defer rows.Close()

	result := make([]dto.NameSpace, 0)

	for rows.Next() {
		var namespace dto.NameSpace

		err = rows.Scan(
			&namespace.ID,
			&namespace.Name,
			&namespace.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("scan namespace: %w", err)
		}

		result = append(result, namespace)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate namespaces: %w", err)
	}

	return result, nil
}
