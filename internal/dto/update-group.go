package dto

type UpdateGroup struct {
	ID                int64   `json:"id"`
	Name              *string `json:"name"`
	RollingPercentage *int    `json:"rolling_percentage"`
}
