package dto

import "time"

type LayerExperiment struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	RolloutPercentage int64     `json:"rollout_percentage"`
	Buckets           []int64   `json:"buckets"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	Status            string    `json:"status"`
}
