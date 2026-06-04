package namespace

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
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
		return nil, shortcut.MapStorageError(err)
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
			return nil, shortcut.MapStorageError(err)
		}

		result = append(result, namespace)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return result, nil
}
