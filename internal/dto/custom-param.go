package dto

type CustomParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NameSpaceID int64  `json:"namespace_id"`
	Type        string `json:"type"`
}
