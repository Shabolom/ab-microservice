package dto

import "time"

type RawExperiment struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	NameSpaceName     string `json:"name_space_name"`
	RollingPercentage int    `json:"rolling_percentage"`
	StartDate         time.Time
	EndDate           time.Time
	NameSpace         NameSpace
	Layer             []Layer
	Group             []Group
}
