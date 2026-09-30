package customParams

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetList(ctx context.Context) ([]*dto.CustomParams, error) {
	query := `
		SELECT
			id,
			name,
			namespace_id,
			type
		FROM customparameter
		ORDER BY id
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}
	defer rows.Close()

	var params []*dto.CustomParams

	for rows.Next() {
		var param dto.CustomParams

		err = rows.Scan(
			&param.ID,
			&param.Name,
			&param.NameSpaceID,
			&param.Type,
		)
		if err != nil {
			return nil, shortcut.MapStorageError(err)
		}

		params = append(params, &param)
	}

	if err = rows.Err(); err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return params, nil
}
