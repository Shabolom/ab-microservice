package dto

type LayerBuckets struct {
	LayerID int64   `json:"layer_id"`
	Buckets []int64 `json:"buckets"`
}
