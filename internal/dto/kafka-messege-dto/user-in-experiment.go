package kafkaMessageDto

import "time"

type UserInExperimentMessage struct {
	UserID         int64     `json:"user_id"`
	ExperimentName string    `json:"experiment_name"`
	GroupName      string    `json:"group_name"`
	Timestamp      time.Time `json:"timestamp"`
}
