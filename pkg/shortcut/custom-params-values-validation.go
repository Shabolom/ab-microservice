package shortcut

const (
	ParamTypeString      = "STRING"
	ParamTypeInt         = "INT"
	ParamTypeSemver      = "SEMVER"
	ParamTypeDate        = "DATE"
	ParamTypeStringArray = "STRING_ARRAY"
	ParamTypeIntArray    = "INT_ARRAY"
	ParamTypeSemverArray = "SEMVER_ARRAY"
	ParamTypeDateArray   = "DATE_ARRAY"
)

var AllowedConditions = map[string]map[string]struct{}{
	ParamTypeString: {
		"=":            {},
		"<>":           {},
		"CONTAINS":     {},
		"NOT CONTAINS": {},
	},

	ParamTypeInt: {
		"=":           {},
		"<>":          {},
		">":           {},
		"<":           {},
		">=":          {},
		"<=":          {},
		"BETWEEN":     {},
		"NOT BETWEEN": {},
	},

	ParamTypeDate: {
		"=":           {},
		"<>":          {},
		">":           {},
		"<":           {},
		">=":          {},
		"<=":          {},
		"BETWEEN":     {},
		"NOT BETWEEN": {},
	},

	ParamTypeSemver: {
		"=":           {},
		"<>":          {},
		">":           {},
		"<":           {},
		">=":          {},
		"<=":          {},
		"BETWEEN":     {},
		"NOT BETWEEN": {},
	},

	ParamTypeStringArray: {
		"IN":         {},
		"NOT IN":     {},
		"ONE OF":     {},
		"NOT ONE OF": {},
	},

	ParamTypeIntArray: {
		"IN":         {},
		"NOT IN":     {},
		"ONE OF":     {},
		"NOT ONE OF": {},
	},

	ParamTypeDateArray: {
		"IN":         {},
		"NOT IN":     {},
		"ONE OF":     {},
		"NOT ONE OF": {},
	},

	ParamTypeSemverArray: {
		"IN":         {},
		"NOT IN":     {},
		"ONE OF":     {},
		"NOT ONE OF": {},
	},
}
