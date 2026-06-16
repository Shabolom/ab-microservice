package dto

type Group struct {
	ID                int64   `json:"id,omitempty"`
	Name              string  `json:"name,omitempty"`
	RollingPercentage int64   `json:"rolling_percentage,omitempty"`
	DeviceID          []int64 `json:"device_id,omitempty"`
}
