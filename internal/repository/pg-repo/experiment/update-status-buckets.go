package experiment

import "context"

func (s *Storage) UpdateStatusBuckets(ctx context.Context, experimentID int64, status string, buckets []int64) error {
	query := `
		UPDATE experiments
		SET
			status = $2,
			bucket = $3
		WHERE id = $1
	`

	_, err := s.conn.Exec(
		ctx,
		query,
		experimentID,
		status,
		buckets,
	)

	return err
}
