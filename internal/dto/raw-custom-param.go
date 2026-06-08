package dto

type RawCustomParamWithCondition struct {
	ID               int64  `json:"id"`
	ParameterID      int64  `json:"parameter_id"`
	ParameterGroupID int64  `json:"parameter_group_id"`
	ParameterType    string `json:"parameter_type"`
	ParameterName    string `json:"parameter_name"`
	Value            string `json:"value"`
	Condition        string `json:"condition"`
}
