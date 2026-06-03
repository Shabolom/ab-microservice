package dto

type LayerWithExperiments struct {
	LayerID     int64        `json:"layer_id"`
	Experiments []Experiment `json:"experiments"`
}
