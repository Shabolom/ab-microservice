package kafkaMessageDto

type stopedExperiment struct {
	ExpID   int64  `json:"exp_id"`
	ExpName string `json:"exp_name"`
}
