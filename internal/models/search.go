package models

type SearchResp struct {
	Count int64          `json:"count"`
	List  []SearchObject `json:"list"`
}

type SearchObject struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Title          string         `json:"title"`
	Excerpt        string         `json:"excerpt"`
	VoteCount      int            `json:"vote_count"`
	AnswerCount    int            `json:"answer_count"`
	Accepted       bool           `json:"accepted"`
	CreatedAt      int64          `json:"created_at"`
	UserInfo       *UserBasicInfo `json:"user_info"`
	Tags           []TagItem      `json:"tags"`
	QuestionID     string         `json:"question_id,omitempty"`
}

type SearchReq struct {
	Query string `json:"q"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Order string `json:"order,omitempty"`
}
