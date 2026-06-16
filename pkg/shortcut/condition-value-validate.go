package shortcut

func conditionValueValidation(condition string, conditionValue string, reqValue string, paramType string) (bool, error) {
	switch paramType {
	case ParamTypeString:
		return compareString(condition, conditionValue, reqValue)

	case ParamTypeInt:
		return compareInt(condition, conditionValue, reqValue)

	case ParamTypeDate:
		return compareDate(condition, conditionValue, reqValue)

	case ParamTypeSemver:
		return compareSemver(condition, conditionValue, reqValue)

	case ParamTypeStringArray:
		return compareStringArray(condition, conditionValue, reqValue)

	case ParamTypeIntArray:
		return compareIntArray(condition, conditionValue, reqValue)

	case ParamTypeDateArray:
		return compareDateArray(condition, conditionValue, reqValue)

	case ParamTypeSemverArray:
		return compareSemverArray(condition, conditionValue, reqValue)

	default:
		return false, ErrValidation
	}
}
