package shortcut

import (
	"strconv"
)

func compareInt(condition string, conditionValue string, reqValue string) (bool, error) {
	reqInt, err := strconv.ParseInt(reqValue, 10, 64)
	if err != nil {
		return false, ErrValidation
	}

	switch condition {
	case "=", "<>", ">", "<", ">=", "<=":
		conditionInt, err := strconv.ParseInt(conditionValue, 10, 64)
		if err != nil {
			return false, ErrValidation
		}

		switch condition {
		case "=":
			return reqInt == conditionInt, nil
		case "<>":
			return reqInt != conditionInt, nil
		case ">":
			return reqInt > conditionInt, nil
		case "<":
			return reqInt < conditionInt, nil
		case ">=":
			return reqInt >= conditionInt, nil
		default:
			return reqInt <= conditionInt, nil
		}

	case "BETWEEN", "NOT BETWEEN":
		values, err := parseIntArray(conditionValue)
		if err != nil || len(values) != 2 {
			return false, ErrValidation
		}

		from := values[0]
		to := values[1]

		if from > to {
			return false, ErrValidation
		}

		if condition == "BETWEEN" {
			return reqInt >= from && reqInt <= to, nil
		}

		return reqInt < from || reqInt > to, nil
	}

	return false, ErrValidation
}

func compareIntArray(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionValues, err := parseIntArray(conditionValue)
	if err != nil {
		return false, ErrValidation
	}

	reqValues, err := parseIntArray(reqValue)
	if err != nil {
		return false, ErrValidation
	}

	if len(conditionValues) == 0 || len(reqValues) == 0 {
		return false, ErrValidation
	}

	switch condition {
	case "IN":
		return containsAllInt64(conditionValues, reqValues), nil

	case "NOT IN":
		return !containsAllInt64(conditionValues, reqValues), nil

	case "ONE OF":
		return containsAnyInt64(conditionValues, reqValues), nil

	case "NOT ONE OF":
		return !containsAnyInt64(conditionValues, reqValues), nil

	default:
		return false, ErrValidation
	}
}
