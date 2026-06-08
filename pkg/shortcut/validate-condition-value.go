package shortcut

import (
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
)

func ValidateConditionValue(paramType string, condition string, value string) error {
	if err := validateCondition(paramType, condition); err != nil {
		return err
	}

	switch paramType {
	case ParamTypeString:
		if strings.TrimSpace(value) == "" {
			return ErrValidation
		}

	case ParamTypeInt:
		if _, err := strconv.ParseInt(value, 10, 64); err != nil {
			return ErrValidation
		}

	case ParamTypeDate:
		if _, err := time.Parse(time.RFC3339, value); err != nil {
			return ErrValidation
		}

	case ParamTypeSemver:
		if _, err := semver.NewVersion(value); err != nil {
			return ErrValidation
		}

	case ParamTypeStringArray:
		if len(parseStringArray(value)) == 0 {
			return ErrValidation
		}

	case ParamTypeIntArray:
		if _, err := parseIntArray(value); err != nil {
			return ErrValidation
		}

	case ParamTypeDateArray:
		if _, err := parseDateArray(value); err != nil {
			return ErrValidation
		}

	case ParamTypeSemverArray:
		if _, err := parseSemverArray(value); err != nil {
			return ErrValidation
		}

	default:
		return ErrValidation
	}

	return nil
}

func validateCondition(paramType string, condition string) error {
	paramType = strings.ToUpper(paramType)

	conditions, ok := AllowedConditions[paramType]
	if !ok {
		return ErrNotSupportedType
	}

	if _, ok := conditions[condition]; !ok {
		return ErrNotSupportedConditions
	}

	return nil
}

func parseStringArray(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

func parseIntArray(value string) ([]int64, error) {
	parts := strings.Split(value, ",")
	result := make([]int64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, ErrValidation
		}

		v, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}

func parseDateArray(value string) ([]time.Time, error) {
	parts := strings.Split(value, ",")
	result := make([]time.Time, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, ErrValidation
		}

		v, err := time.Parse(time.RFC3339, part)
		if err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}

func parseSemverArray(value string) ([]*semver.Version, error) {
	parts := strings.Split(value, ",")
	result := make([]*semver.Version, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, ErrValidation
		}

		v, err := semver.NewVersion(part)
		if err != nil {
			return nil, err
		}

		result = append(result, v)
	}

	return result, nil
}
