package customParams

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) Create(ctx context.Context, param *dto.CustomParams) (*dto.CustomParams, error) {
	query := `
		INSERT INTO customparameter (
			name,
			namespace_id,
			type
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			name,
			namespace_id,
			type
	`

	created := &dto.CustomParams{}

	err := s.conn.QueryRow(
		ctx,
		query,
		param.Name,
		param.NameSpaceID,
		param.Type,
	).Scan(
		&created.ID,
		&created.Name,
		&created.NameSpaceID,
		&created.Type,
	)
	if err != nil {
		return nil, shortcut.MapStorageError(err)
	}

	return created, nil
}
