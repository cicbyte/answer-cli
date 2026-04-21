package models

import "time"

type QuestionPageResp struct {
	Count int64               `json:"count"`
	List  []QuestionListItem  `json:"list"`
}

type QuestionListItem struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	ViewCount       int    `json:"view_count"`
	VoteCount       int    `json:"vote_count"`
	AnswerCount     int    `json:"answer_count"`
	CollectionCount int    `json:"collection_count"`
	CommentCount    int    `json:"comment_count"`
	Status          int    `json:"status"`
	Pin             int    `json:"pin"`
	AcceptedAnswerID string `json:"accepted_answer_id"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
	UserInfo        *UserBasicInfo `json:"user_info"`
	Tags            []TagItem     `json:"tags"`
}

type QuestionInfoResp struct {
	ID               string         `json:"id"`
	Title            string         `json:"title"`
	Content          string         `json:"content"`
	HTML             string         `json:"html"`
	ViewCount        int            `json:"view_count"`
	UniqueViewCount  int            `json:"unique_view_count"`
	VoteCount        int            `json:"vote_count"`
	AnswerCount      int            `json:"answer_count"`
	CollectionCount  int            `json:"collection_count"`
	FollowCount      int            `json:"follow_count"`
	CommentCount     int            `json:"comment_count"`
	Status           int            `json:"status"`
	Pin              int            `json:"pin"`
	AcceptedAnswerID string         `json:"accepted_answer_id"`
	LastAnswerID     string         `json:"last_answer_id"`
	HotScore         int            `json:"hot_score"`
	CreatedAt        int64          `json:"created_at"`
	UpdatedAt        int64          `json:"updated_at"`
	UserInfo         *UserBasicInfo `json:"user_info"`
	Tags             []TagItem      `json:"tags"`
}

type QuestionAddReq struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags,omitempty"`
}

type QuestionUpdateReq struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
}

type QuestionDeleteReq struct {
	ID string `json:"id"`
}

type QuestionCloseReq struct {
	ID       string `json:"id"`
	CloseType int   `json:"close_type"`
	CloseMsg  string `json:"close_msg,omitempty"`
}

type QuestionReopenReq struct {
	ID string `json:"id"`
}

type QuestionListReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Order    string `json:"order,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Username string `json:"username,omitempty"`
	InDays   int    `json:"in_days,omitempty"`
}

// helper
func FormatTimestamp(ts int64) time.Time {
	if ts == 0 {
		return time.Time{}
	}
	sec := ts
	if ts > 1e12 {
		sec = ts / 1000
	}
	return time.Unix(sec, 0)
}
