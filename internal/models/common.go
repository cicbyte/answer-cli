package models

type PageModel struct {
	Count int64       `json:"count"`
	List  any `json:"list"`
}

type TagItem struct {
	ID          string `json:"id"`
	SlugName    string `json:"slug_name"`
	DisplayName string `json:"display_name"`
	MainTagID   int64  `json:"main_tag_id"`
	Reserved    bool   `json:"reserved"`
}

type UserBasicInfo struct {
	ID          string      `json:"id"`
	Username    string      `json:"username"`
	DisplayName string      `json:"display_name"`
	Avatar      *AvatarInfo `json:"avatar"`
	Rank        int         `json:"rank"`
}
