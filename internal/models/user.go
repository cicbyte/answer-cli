package models

// UserDetailResp 用户详情
type UserDetailResp struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
	Rank        int    `json:"rank"`
	Bio         string `json:"bio"`
	Website     string `json:"website"`
	Location    string `json:"location"`
}
