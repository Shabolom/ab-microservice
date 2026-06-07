package dto

import "time"

type Experiment struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	NameSpace         string    `json:"namespace"`
	Status            string    `json:"status"`
	RolloutPercentage int64     `json:"rollout_percentage"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	PassingCities     []string  `json:"passing_cities"`
	ExcludedCities    []string  `json:"excluded_cities"`
	PassingStores     []string  `json:"passing_stores"`
	ExcludedStores    []string  `json:"excluded_stores"`
	LayersID          []int64   `json:"layers_id"`
	Groups            []Group   `json:"groups"`
}
