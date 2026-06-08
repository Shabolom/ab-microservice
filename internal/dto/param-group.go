package dto

type ParamGroup struct {
	ID                   int64                      `json:"id"`
	Percent              int64                      `json:"percent"`
	ParamsWithConditions []CustomParamWithCondition `json:"paramswithconditions"`
}

type ParamGroupWithRawParams struct {
	ID                   int64                         `json:"id"`
	Percent              int64                         `json:"percent"`
	ParamsWithConditions []RawCustomParamWithCondition `json:"paramswithconditions"`
}
