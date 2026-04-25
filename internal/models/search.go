package models

type SearchResp struct {
	Count int64           `json:"count"`
	List  []*SearchResult `json:"list"`
}

type SearchResult struct {
	ObjectType string        `json:"object_type"`
	Object     *SearchObject `json:"object"`
}

type SearchObject struct {
	ID          string `json:"id"`
	QuestionID  string `json:"question_id"`
	Title       string `json:"title"`
	Excerpt     string `json:"excerpt"`
	VoteCount   int    `json:"vote_count"`
	AnswerCount int    `json:"answer_count"`
	Accepted    bool   `json:"accepted"`
	CreatedAt   int64  `json:"created_at"`
	StatusStr   string `json:"status"`
}

type SearchReq struct {
	Query string `json:"q"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Order string `json:"order,omitempty"`
}
