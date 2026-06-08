package shortcut

import (
	"time"

	"github.com/Masterminds/semver/v3"
)

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func containsInt64(values []int64, target int64) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func containsAllStrings(source []string, targets []string) bool {
	for _, target := range targets {
		if !containsString(source, target) {
			return false
		}
	}
	return true
}

func containsAnyString(source []string, targets []string) bool {
	for _, target := range targets {
		if containsString(source, target) {
			return true
		}
	}
	return false
}

func containsAllInt64(source []int64, targets []int64) bool {
	for _, target := range targets {
		if !containsInt64(source, target) {
			return false
		}
	}
	return true
}

func containsAnyInt64(source []int64, targets []int64) bool {
	for _, target := range targets {
		if containsInt64(source, target) {
			return true
		}
	}
	return false
}

func containsAllDates(source []time.Time, targets []time.Time) bool {
	for _, target := range targets {
		found := false

		for _, value := range source {
			if value.Equal(target) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func containsAnyDate(source []time.Time, targets []time.Time) bool {
	for _, target := range targets {
		for _, value := range source {
			if value.Equal(target) {
				return true
			}
		}
	}

	return false
}

func containsAllSemver(source []*semver.Version, targets []*semver.Version) bool {
	for _, target := range targets {
		found := false

		for _, value := range source {
			if value.Equal(target) {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func containsAnySemver(source []*semver.Version, targets []*semver.Version) bool {
	for _, target := range targets {
		for _, value := range source {
			if value.Equal(target) {
				return true
			}
		}
	}

	return false
}
