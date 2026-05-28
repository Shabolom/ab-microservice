package dto

type Group struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty" json:"name,omitempty"`
	RollingPercentage int    `json:"rolling_percentage,omitempty" json:"rolling_percentage,omitempty"`
}
