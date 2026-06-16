package shortcut

import (
	"time"
)

func compareDate(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionDate, err := time.Parse(time.RFC3339, conditionValue)
	if err != nil {
		return false, ErrValidation
	}

	reqDate, err := time.Parse(time.RFC3339, reqValue)
	if err != nil {
		return false, ErrValidation
	}

	switch condition {
	case "=":
		return reqDate.Equal(conditionDate), nil
	case "<>":
		return !reqDate.Equal(conditionDate), nil
	case ">":
		return reqDate.After(conditionDate), nil
	case "<":
		return reqDate.Before(conditionDate), nil
	case ">=":
		return reqDate.After(conditionDate) || reqDate.Equal(conditionDate), nil
	case "<=":
		return reqDate.Before(conditionDate) || reqDate.Equal(conditionDate), nil
	case "BETWEEN":
		values, err := parseDateArray(conditionValue)
		if err != nil || len(values) != 2 {
			return false, ErrValidation
		}
		return (reqDate.After(values[0]) || reqDate.Equal(values[0])) &&
			(reqDate.Before(values[1]) || reqDate.Equal(values[1])), nil
	case "NOT BETWEEN":
		values, err := parseDateArray(conditionValue)
		if err != nil || len(values) != 2 {
			return false, ErrValidation
		}
		return reqDate.Before(values[0]) || reqDate.After(values[1]), nil
	default:
		return false, ErrValidation
	}
}

func compareDateArray(condition string, conditionValue string, reqValue string) (bool, error) {
	conditionValues, err := parseDateArray(conditionValue)
	if err != nil {
		return false, ErrValidation
	}

	reqValues, err := parseDateArray(reqValue)
	if err != nil {
		return false, ErrValidation
	}

	switch condition {
	case "CONTAINS":
		return containsAllDates(reqValues, conditionValues), nil
	case "NOT CONTAINS":
		return !containsAllDates(reqValues, conditionValues), nil
	case "ONE OF":
		return containsAnyDate(reqValues, conditionValues), nil
	case "NOT ONE OF":
		return !containsAnyDate(reqValues, conditionValues), nil
	default:
		return false, ErrValidation
	}
}
