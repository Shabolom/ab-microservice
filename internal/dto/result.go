package dto

type GetExperimentsReply struct {
	ExperimentName string `json:"experiment_name,omitempty"`
	GroupName      string `json:"group_name,omitempty"`
}
