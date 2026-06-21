package dto

import "time"

const (
	FeatureToggleStatusDraft    = "draft"
	FeatureToggleStatusActive   = "active"
	FeatureToggleStatusDisabled = "disabled"
	FeatureToggleStatusArchived = "archived"
)

type FeatureToggle struct {
	ID                int64
	RolloutPercentage int
	NamespaceID       int64
	Name              string
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
