package models

import "encoding/json"

type LoginResponse struct {
	AccessToken string          `json:"access_token"`
	ID          string          `json:"id"`
	Username    string          `json:"username"`
	EMail       string          `json:"e_mail"`
	DisplayName string          `json:"display_name"`
	Avatar      json.RawMessage `json:"avatar"`
	Rank        int             `json:"rank"`
}

type AvatarInfo struct {
	Type     string `json:"type"`
	Gravatar string `json:"gravatar"`
	Custom   string `json:"custom"`
}

func (a *AvatarInfo) GetURL() string {
	switch a.Type {
	case "gravatar":
		return a.Gravatar
	case "custom":
		return a.Custom
	default:
		return a.Gravatar
	}
}

type UserInfoResp struct {
	ID            string      `json:"id"`
	Username      string      `json:"username"`
	DisplayName   string      `json:"display_name"`
	EMail         string      `json:"e_mail"`
	Avatar        *AvatarInfo `json:"avatar"`
	Bio           string      `json:"bio"`
	Website       string      `json:"website"`
	Location      string      `json:"location"`
	Rank          int         `json:"rank"`
	AnswerCount   int         `json:"answer_count"`
	QuestionCount int         `json:"question_count"`
	FollowCount   int         `json:"follow_count"`
	IsAdmin       bool        `json:"is_admin"`
}
