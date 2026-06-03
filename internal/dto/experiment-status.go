package dto

import "time"

type ExperimentStatus struct {
	ExpID     int64
	LayerID   int64
	StartedAt time.Time
	EndedAt   time.Time
}
