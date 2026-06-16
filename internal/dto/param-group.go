package dto

type ParamGroup struct {
	ID                   int64                      `json:"id"`
	Percent              int64                      `json:"percent"`
	ParamsWithConditions []CustomParamWithCondition `json:"params_with_conditions"`
}

type ParamGroupWithRawParams struct {
	ID                   int64                         `json:"id"`
	Percent              int64                         `json:"percent"`
	ParamsWithConditions []RawCustomParamWithCondition `json:"params_with_conditions"`
}
