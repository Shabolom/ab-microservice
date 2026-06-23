package dto

type UserInFeatureReq struct {
	UserId    int64  `json:"user_id"`
	Namespace string `json:"namespace"`
	Platform  string `json:"platform"`
}
