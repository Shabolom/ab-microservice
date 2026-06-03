package dto

import "time"

type UpdateExperiment struct {
	ID int64 `json:"id"`

	Name              *string    `json:"name"`
	RolloutPercentage *int64     `json:"rollout_percentage"`
	Bucket            []int64    `json:"bucket"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
	Status            *string    `json:"status"`
}
