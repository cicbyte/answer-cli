package models

type VoteReq struct {
	ObjectID string `json:"object_id"`
	Group    string `json:"group,omitempty"`
	IsCancel bool   `json:"is_cancel,omitempty"`
}

type VoteResp struct {
	ID       string `json:"id"`
	ObjectID string `json:"object_id"`
	Group    string `json:"group"`
	Status   int    `json:"status"`
	UpDown   int    `json:"up_down"`
}
