package dto

import "time"

type RawFeatureToggle struct {
	ID          int64
	NamespaceID int64
	Name        string
	Status      string

	RolloutPercentage *int64
	IOS               *int64
	Android           *int64
	Web               *int64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
