package shortcut

import (
	"strings"
)

func compareString(condition string, conditionValue string, reqValue string) (bool, error) {
	switch condition {
	case "=":
		return reqValue == conditionValue, nil

	case "<>":
		return reqValue != conditionValue, nil

	case "CONTAINS":
		return strings.Contains(reqValue, conditionValue), nil

	case "NOT CONTAINS":
		return !strings.Contains(reqValue, conditionValue), nil

	default:
		return false, ErrValidation
	}
}

func compareStringArray(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionValues := parseStringArray(conditionValue)
	reqValues := parseStringArray(reqValue)

	if len(conditionValues) == 0 || len(reqValues) == 0 {
		return false, ErrValidation
	}

	switch condition {
	case "IN":
		return containsAllStrings(conditionValues, reqValues), nil

	case "NOT IN":
		return !containsAllStrings(conditionValues, reqValues), nil

	case "ONE OF":
		return containsAnyString(conditionValues, reqValues), nil

	case "NOT ONE OF":
		return !containsAnyString(conditionValues, reqValues), nil

	default:
		return false, ErrValidation
	}
}
