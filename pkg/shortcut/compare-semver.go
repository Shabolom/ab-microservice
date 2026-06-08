package shortcut

import (
	"github.com/Masterminds/semver/v3"
)

func compareSemver(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionVersion, err := semver.NewVersion(conditionValue)
	if err != nil {
		return false, ErrValidation
	}

	reqVersion, err := semver.NewVersion(reqValue)
	if err != nil {
		return false, ErrValidation
	}

	switch condition {
	case "=":
		return reqVersion.Equal(conditionVersion), nil
	case "<>":
		return !reqVersion.Equal(conditionVersion), nil
	case ">":
		return reqVersion.GreaterThan(conditionVersion), nil
	case "<":
		return reqVersion.LessThan(conditionVersion), nil
	case ">=":
		return reqVersion.GreaterThan(conditionVersion) || reqVersion.Equal(conditionVersion), nil
	case "<=":
		return reqVersion.LessThan(conditionVersion) || reqVersion.Equal(conditionVersion), nil
	case "BETWEEN":
		values, err := parseSemverArray(conditionValue)
		if err != nil || len(values) != 2 {
			return false, ErrValidation
		}
		return (reqVersion.GreaterThan(values[0]) || reqVersion.Equal(values[0])) &&
			(reqVersion.LessThan(values[1]) || reqVersion.Equal(values[1])), nil
	case "NOT BETWEEN":
		values, err := parseSemverArray(conditionValue)
		if err != nil || len(values) != 2 {
			return false, ErrValidation
		}
		return reqVersion.LessThan(values[0]) || reqVersion.GreaterThan(values[1]), nil
	default:
		return false, ErrValidation
	}
}

func compareSemverArray(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionValues, err := parseSemverArray(conditionValue)
	if err != nil {
		return false, ErrValidation
	}

	reqValues, err := parseSemverArray(reqValue)
	if err != nil {
		return false, ErrValidation
	}

	switch condition {
	case "CONTAINS":
		return containsAllSemver(reqValues, conditionValues), nil
	case "NOT CONTAINS":
		return !containsAllSemver(reqValues, conditionValues), nil
	case "ONE OF":
		return containsAnySemver(reqValues, conditionValues), nil
	case "NOT ONE OF":
		return !containsAnySemver(reqValues, conditionValues), nil
	default:
		return false, ErrValidation
	}
}
