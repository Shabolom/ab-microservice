package dto

type UpdateExperiment struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
	//Namespace         *string `json:"namespace"`
	RollingPercentage *int `json:"rolling_percentage"`
}
