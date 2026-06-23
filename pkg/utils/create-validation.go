package utils

import (
	"ab/internal/dto"
	"ab/pkg/shortcut"
)

func CreateValidation(featureToggle *dto.FeatureToggle) error {
	switch {
	case featureToggle.Name == "":
		return shortcut.ErrFeatureToggleNameRequired
	case featureToggle.NamespaceID <= 0:
		return shortcut.ErrFeatureToggleNamespaceIDRequired
	}

	if err := validateRolloutPercentage(featureToggle.RolloutPercentage); err != nil {
		return err
	}

	if err := validateRolloutPercentage(featureToggle.IOS); err != nil {
		return err
	}

	if err := validateRolloutPercentage(featureToggle.Android); err != nil {
		return err
	}

	if err := validateRolloutPercentage(featureToggle.Web); err != nil {
		return err
	}

	return nil
}

func validateRolloutPercentage(percentage *int64) error {
	if percentage == nil {
		return nil
	}

	if *percentage < 0 || *percentage > 100 {
		return shortcut.ErrFeatureToggleRolloutOutOfRange
	}

	return nil
}
