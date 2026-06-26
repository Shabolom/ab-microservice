package dto

type FeatureTogglePercentageUpdate struct {
	FeatureID  int64  `json:"feature_id"`
	Percentage *int64 `json:"percentage"`
	Ios        *int64 `json:"ios"`
	Android    *int64 `json:"android"`
	Web        *int64 `json:"web"`
}
