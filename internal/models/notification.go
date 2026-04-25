package models

type NotificationListResp struct {
	Count int64              `json:"count"`
	List  []NotificationItem `json:"list"`
}

type NotificationItem struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	IsRead    bool   `json:"is_read"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ObjectID  string `json:"object_id"`
	Reason    string `json:"reason"`
	CreatedAt int64  `json:"created_at"`
}

type NotificationStatusResp struct {
	Inbox       bool `json:"inbox"`
	Achievement bool `json:"achievement"`
}

type NotificationReadAllReq struct {
	Type string `json:"type,omitempty"`
}

type NotificationReadReq struct {
	ID string `json:"id"`
}

type NotificationListReq struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Type      string `json:"type,omitempty"`
	InboxType string `json:"inbox_type,omitempty"`
}
