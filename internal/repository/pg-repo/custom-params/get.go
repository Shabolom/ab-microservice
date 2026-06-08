package customParams

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
	"context"
)

func (s *Storage) GetById(ctx context.Context, id int64) (dto.CustomParams, error) {
	query := `
		SELECT
			id,
			name,
			namespace_id,
			type
		FROM config_app_customparameter
		WHERE id = $1
	`

	var param dto.CustomParams

	err := s.conn.QueryRow(ctx, query, id).Scan(
		&param.ID,
		&param.Name,
		&param.NameSpaceID,
		&param.Type,
	)
	if err != nil {
		return dto.CustomParams{}, shortcut.MapStorageError(err)
	}

	return param, nil
}
