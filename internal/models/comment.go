package models

type CommentListResp struct {
	Count int64           `json:"count"`
	List  []CommentInfo   `json:"list"`
}

type CommentInfo struct {
	ID           string         `json:"id"`
	ObjectID     string         `json:"object_id"`
	QuestionID   string         `json:"question_id"`
	ReplyUserID  string         `json:"reply_user_id,omitempty"`
	ReplyCommentID string       `json:"reply_comment_id,omitempty"`
	OriginalText string         `json:"original_text"`
	ParsedText   string         `json:"parsed_text"`
	VoteCount    int            `json:"vote_count"`
	Status       int            `json:"status"`
	CreatedAt    int64          `json:"created_at"`
	UpdatedAt    int64          `json:"updated_at"`
	UserInfo     *UserBasicInfo `json:"user_info"`
}

type CommentAddReq struct {
	ObjectID       string `json:"object_id"`
	OriginalText   string `json:"original_text"`
	ReplyCommentID string `json:"reply_comment_id,omitempty"`
}

type CommentUpdateReq struct {
	ID           string `json:"id"`
	OriginalText string `json:"original_text"`
}

type CommentDeleteReq struct {
	ID string `json:"id"`
}

type CommentListReq struct {
	ObjectID string `json:"object_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
