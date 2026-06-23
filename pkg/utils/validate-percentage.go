package utils

import "ab/pkg/shortcut"

func ValidatePercentage(percentage *int64) error {
	if percentage == nil {
		return nil
	}

	if *percentage < 0 || *percentage > 100 {
		return shortcut.ErrFeatureToggleRolloutOutOfRange
	}

	return nil
}
