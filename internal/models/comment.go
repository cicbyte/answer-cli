package models

type CommentListResp struct {
	Count int64         `json:"count"`
	List  []CommentInfo `json:"list"`
}

type CommentInfo struct {
	CommentID            string `json:"comment_id"`
	ObjectID             string `json:"object_id"`
	QuestionID           string `json:"question_id"`
	ReplyUserID          string `json:"reply_user_id"`
	ReplyUsername        string `json:"reply_username"`
	ReplyUserDisplayName string `json:"reply_user_display_name"`
	ReplyCommentID       string `json:"reply_comment_id"`
	OriginalText         string `json:"original_text"`
	ParsedText           string `json:"parsed_text"`
	VoteCount            int    `json:"vote_count"`
	IsVote               bool   `json:"is_vote"`
	Status               int    `json:"status"`
	CreatedAt            int64  `json:"created_at"`
	UpdatedAt            int64  `json:"updated_at"`
	UserID               string `json:"user_id"`
	Username             string `json:"username"`
	UserDisplayName      string `json:"user_display_name"`
	UserAvatar           string `json:"user_avatar"`
	UserStatus           string `json:"user_status"`
}

// DisplayAuthor returns the best available display name.
func (c CommentInfo) DisplayAuthor() string {
	if c.UserDisplayName != "" {
		return c.UserDisplayName
	}
	if c.Username != "" {
		return c.Username
	}
	return "匿名"
}

type CommentAddReq struct {
	ObjectID       string `json:"object_id"`
	OriginalText   string `json:"original_text"`
	ReplyCommentID string `json:"reply_comment_id,omitempty"`
}

type CommentUpdateReq struct {
	CommentID    string `json:"comment_id"`
	OriginalText string `json:"original_text"`
}

type CommentDeleteReq struct {
	CommentID string `json:"comment_id"`
}

type CommentListReq struct {
	ObjectID string `json:"object_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
