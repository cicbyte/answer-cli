package models

type TagListResp struct {
	Count int64       `json:"count"`
	List  []TagDetail `json:"list"`
}

type TagDetail struct {
	ID           string `json:"id"`
	SlugName     string `json:"slug_name"`
	DisplayName  string `json:"display_name"`
	OriginalText string `json:"original_text"`
	ParsedText   string `json:"parsed_text"`
	FollowCount  int    `json:"follow_count"`
	QuestionCount int   `json:"question_count"`
	Status       int    `json:"status"`
	Recommend    bool   `json:"recommend"`
	Reserved     bool   `json:"reserved"`
	MainTagID    int64  `json:"main_tag_id"`
	MainTagSlugName string `json:"main_tag_slug_name"`
	CreatedAt    int64  `json:"created_at"`
}

type TagAddReq struct {
	SlugName     string `json:"slug_name"`
	DisplayName  string `json:"display_name"`
	OriginalText string `json:"original_text,omitempty"`
}

type TagUpdateReq struct {
	SlugName     string `json:"slug_name"`
	DisplayName  string `json:"display_name"`
	OriginalText string `json:"original_text,omitempty"`
}

type TagDeleteReq struct {
	SlugName string `json:"slug_name"`
}

type TagListReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	QueryCond string `json:"query_cond,omitempty"`
}
