package dto

import "time"

type RawExperiment struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	NameSpaceName     string `json:"name_space_name"`
	RollingPercentage int64  `json:"rolling_percentage"`
	Status            string `json:"status"`
	StartDate         time.Time
	EndDate           time.Time
	NameSpace         NameSpace
	Bucket            []int64
	Layer             []Layer
	Group             []Group
}
