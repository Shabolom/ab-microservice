package dto

import "time"

type RawExperiment struct {
	Id                 int64                     `json:"id"`
	Name               string                    `json:"name"`
	NamespaceName      string                    `json:"namespace_name"`
	RollingPercentage  int64                     `json:"rolling_percentage"`
	Status             string                    `json:"status"`
	StartDate          time.Time                 `json:"start_date"`
	EndDate            time.Time                 `json:"end_date"`
	NameSpace          NameSpace                 `json:"name_space"`
	PassingCities      []string                  `json:"passing_cities"`
	ExcludedCities     []string                  `json:"excluded_cities"`
	PassingStores      []string                  `json:"passing_stores"`
	ExcludedStores     []string                  `json:"excluded_stores"`
	Bucket             []int64                   `json:"bucket"`
	Layer              []Layer                   `json:"layer"`
	CustomParamsGroups []ParamGroupWithRawParams `json:"custom_params_groups"`
	Group              []Group                   `json:"group"`
}
