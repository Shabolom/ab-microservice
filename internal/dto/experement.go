package dto

type Experiment struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	NameSpase         string `json:"name_spase"`
	RollingPercentage int    `json:"rolling_percentage"`
}
