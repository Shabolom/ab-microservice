package dto

import "time"

type Experiment struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	Status            string    `json:"status"`
	RolloutPercentage int64     `json:"rollout_percentage"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	LayersID          []int64   `json:"layers_id"`
	Groups            []Group   `json:"groups"`
}
