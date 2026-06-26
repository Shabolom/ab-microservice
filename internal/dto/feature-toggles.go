package dto

import "time"

const (
	FeatureToggleStatusDraft     = "draft"
	FeatureToggleStatusActive    = "active"
	FeatureToggleStatusDisabled  = "disabled"
	FeatureToggleStatusArchived  = "archived"
	FeatureToggleWebPlatform     = "web"
	FeatureToggleIosPlatform     = "ios"
	FeatureToggleAndroidPlatform = "android"
)

type FeatureToggle struct {
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
