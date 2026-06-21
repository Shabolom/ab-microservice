package dto

import "time"

type RawFeatureToggle struct {
	ID                int64
	RolloutPercentage int64
	Buckets           []int64
	NamespaceID       int64
	Name              string
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
