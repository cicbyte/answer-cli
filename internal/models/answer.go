package models

type AnswerListResp struct {
	Count int64        `json:"count"`
	List  []AnswerInfo `json:"list"`
}

type AnswerInfo struct {
	ID           string         `json:"id"`
	QuestionID   string         `json:"question_id"`
	Content      string         `json:"content"`
	HTML         string         `json:"html"`
	VoteCount    int            `json:"vote_count"`
	CommentCount int            `json:"comment_count"`
	Accepted     int            `json:"accepted"`
	Status       int            `json:"status"`
	CreatedAt    int64          `json:"create_time"`
	UpdatedAt    int64          `json:"update_time"`
	UserInfo     *UserBasicInfo `json:"user_info"`
}

type AnswerAddReq struct {
	QuestionID string `json:"question_id"`
	Content    string `json:"content"`
}

type AnswerUpdateReq struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type AnswerDeleteReq struct {
	ID string `json:"id"`
}

type AnswerAcceptReq struct {
	QuestionID string `json:"question_id"`
	AnswerID   string `json:"answer_id"`
}

type AnswerListReq struct {
	QuestionID string `json:"question_id"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	Order      string `json:"order,omitempty"`
}
