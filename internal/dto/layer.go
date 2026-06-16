package dto

type Layer struct {
	ID          int64  `json:"id"`
	NameSpaceID int64  `json:"namespace_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
