package dto

type Group struct {
	ID                int64  `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	RollingPercentage int    `json:"rolling_percentage,omitempty"`
}
