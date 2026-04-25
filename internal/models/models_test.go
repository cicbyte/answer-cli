package models

import (
	"encoding/json"
	"testing"
	"time"
)

// --- FormatTimestamp ---

func TestFormatTimestamp_Zero(t *testing.T) {
	ts := FormatTimestamp(0)
	if !ts.IsZero() {
		t.Fatalf("expected zero time, got %v", ts)
	}
}

func TestFormatTimestamp_Seconds(t *testing.T) {
	ts := FormatTimestamp(1714560000)
	expected := time.Unix(1714560000, 0)
	if !ts.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, ts)
	}
}

func TestFormatTimestamp_Milliseconds(t *testing.T) {
	// >1e12 判定为毫秒，应除以 1000
	ts := FormatTimestamp(1714560000000)
	expected := time.Unix(1714560000, 0)
	if !ts.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, ts)
	}
}

func TestFormatTimestamp_Negative(t *testing.T) {
	ts := FormatTimestamp(-1)
	// -1 是有效的 Unix 秒时间戳（1969-12-31 23:59:59 UTC）
	if ts.IsZero() {
		t.Fatal("expected non-zero time for -1")
	}
	expected := time.Unix(-1, 0)
	if !ts.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, ts)
	}
}

// --- CommentInfo.DisplayAuthor ---

func TestDisplayAuthor_DisplayName(t *testing.T) {
	c := CommentInfo{UserDisplayName: "张三", Username: "zhangsan"}
	if got := c.DisplayAuthor(); got != "张三" {
		t.Fatalf("expected 张三, got %s", got)
	}
}

func TestDisplayAuthor_Username(t *testing.T) {
	c := CommentInfo{Username: "zhangsan"}
	if got := c.DisplayAuthor(); got != "zhangsan" {
		t.Fatalf("expected zhangsan, got %s", got)
	}
}

func TestDisplayAuthor_Anonymous(t *testing.T) {
	c := CommentInfo{}
	if got := c.DisplayAuthor(); got != "匿名" {
		t.Fatalf("expected 匿名, got %s", got)
	}
}

func TestDisplayAuthor_EmptyDisplayName(t *testing.T) {
	c := CommentInfo{UserDisplayName: "", Username: "admin"}
	if got := c.DisplayAuthor(); got != "admin" {
		t.Fatalf("expected admin, got %s", got)
	}
}

// --- QuestionListItem.DisplayAuthor ---

func TestQuestionDisplayAuthor_Operator(t *testing.T) {
	q := QuestionListItem{Operator: &UserBasicInfo{DisplayName: "李四", Username: "lisi"}}
	if got := q.DisplayAuthor(); got != "李四" {
		t.Fatalf("expected 李四, got %s", got)
	}
}

func TestQuestionDisplayAuthor_OperatorUsername(t *testing.T) {
	q := QuestionListItem{Operator: &UserBasicInfo{DisplayName: "", Username: "lisi"}}
	if got := q.DisplayAuthor(); got != "lisi" {
		t.Fatalf("expected lisi, got %s", got)
	}
}

func TestQuestionDisplayAuthor_NilOperator(t *testing.T) {
	q := QuestionListItem{}
	if got := q.DisplayAuthor(); got != "匿名" {
		t.Fatalf("expected 匿名, got %s", got)
	}
}

// --- AvatarInfo ---

func TestAvatarInfo_UnmarshalJSON_String(t *testing.T) {
	data := `"https://example.com/avatar.png"`
	var a AvatarInfo
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Fatal(err)
	}
	if a.Type != "custom" {
		t.Fatalf("expected type custom, got %s", a.Type)
	}
	if a.Custom != "https://example.com/avatar.png" {
		t.Fatalf("expected custom URL, got %s", a.Custom)
	}
}

func TestAvatarInfo_UnmarshalJSON_Object(t *testing.T) {
	data := `{"type":"gravatar","gravatar":"https://gravatar.com/xxx","custom":""}`
	var a AvatarInfo
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Fatal(err)
	}
	if a.Type != "gravatar" {
		t.Fatalf("expected type gravatar, got %s", a.Type)
	}
	if a.Gravatar != "https://gravatar.com/xxx" {
		t.Fatalf("expected gravatar URL, got %s", a.Gravatar)
	}
}

func TestAvatarInfo_GetURL_Gravatar(t *testing.T) {
	a := AvatarInfo{Type: "gravatar", Gravatar: "https://gravatar.com/xxx"}
	if got := a.GetURL(); got != "https://gravatar.com/xxx" {
		t.Fatalf("expected gravatar URL, got %s", got)
	}
}

func TestAvatarInfo_GetURL_Custom(t *testing.T) {
	a := AvatarInfo{Type: "custom", Custom: "https://example.com/a.png"}
	if got := a.GetURL(); got != "https://example.com/a.png" {
		t.Fatalf("expected custom URL, got %s", got)
	}
}

func TestAvatarInfo_GetURL_Empty(t *testing.T) {
	a := AvatarInfo{}
	// 默认返回 Gravatar（空字符串）
	if got := a.GetURL(); got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}

// --- JSON 反序列化 ---

func TestCommentInfo_UnmarshalJSON(t *testing.T) {
	raw := `{
		"comment_id": "c123",
		"object_id": "q1",
		"original_text": "hello world",
		"vote_count": 5,
		"created_at": 1714560000,
		"user_id": "u1",
		"username": "admin",
		"user_display_name": "Admin"
	}`
	var c CommentInfo
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatal(err)
	}
	if c.CommentID != "c123" {
		t.Fatalf("expected c123, got %s", c.CommentID)
	}
	if c.ObjectID != "q1" {
		t.Fatalf("expected q1, got %s", c.ObjectID)
	}
	if c.OriginalText != "hello world" {
		t.Fatalf("expected hello world, got %s", c.OriginalText)
	}
	if c.VoteCount != 5 {
		t.Fatalf("expected 5, got %d", c.VoteCount)
	}
	if c.DisplayAuthor() != "Admin" {
		t.Fatalf("expected Admin, got %s", c.DisplayAuthor())
	}
}

func TestQuestionListItem_UnmarshalJSON(t *testing.T) {
	raw := `{
		"id": "q100",
		"title": "test question",
		"vote_count": 10,
		"answer_count": 2,
		"created_at": 1714560000,
		"operator": {
			"id": "u1",
			"username": "admin",
			"display_name": "Admin",
			"rank": 5
		},
		"tags": [
			{"slug_name": "go", "display_name": "go"}
		]
	}`
	var q QuestionListItem
	if err := json.Unmarshal([]byte(raw), &q); err != nil {
		t.Fatal(err)
	}
	if q.ID != "q100" {
		t.Fatalf("expected q100, got %s", q.ID)
	}
	if q.Title != "test question" {
		t.Fatalf("expected test question, got %s", q.Title)
	}
	if q.Operator == nil || q.Operator.DisplayName != "Admin" {
		t.Fatalf("expected operator Admin")
	}
	if len(q.Tags) != 1 || q.Tags[0].SlugName != "go" {
		t.Fatalf("expected 1 tag 'go'")
	}
	if q.DisplayAuthor() != "Admin" {
		t.Fatalf("expected Admin, got %s", q.DisplayAuthor())
	}
}

func TestSearchResult_UnmarshalJSON(t *testing.T) {
	raw := `{
		"object_type": "question",
		"object": {
			"id": "q1",
			"title": "search result",
			"vote_count": 3,
			"answer_count": 1,
			"created_at": 1714560000
		}
	}`
	var sr SearchResult
	if err := json.Unmarshal([]byte(raw), &sr); err != nil {
		t.Fatal(err)
	}
	if sr.ObjectType != "question" {
		t.Fatalf("expected question, got %s", sr.ObjectType)
	}
	if sr.Object == nil || sr.Object.Title != "search result" {
		t.Fatalf("expected object title")
	}
}

func TestNotificationItem_UnmarshalJSON(t *testing.T) {
	raw := `{
		"id": "n1",
		"type": "inbox",
		"is_read": true,
		"title": "new answer",
		"created_at": 1714560000
	}`
	var n NotificationItem
	if err := json.Unmarshal([]byte(raw), &n); err != nil {
		t.Fatal(err)
	}
	if n.ID != "n1" {
		t.Fatalf("expected n1, got %s", n.ID)
	}
	if !n.IsRead {
		t.Fatal("expected is_read true")
	}
}

// --- AppConfig ---

func TestAppConfig_Getters(t *testing.T) {
	c := AppConfig{}
	c.Server.BaseURL = "http://localhost:8080"
	c.Server.Token = "tok123"
	c.AI.Provider = "openai"
	c.AI.Model = "gpt-4"
	c.Output.Format = "json"

	if c.GetServerURL() != "http://localhost:8080" {
		t.Fatalf("expected server URL")
	}
	if c.GetServerToken() != "tok123" {
		t.Fatalf("expected token")
	}
	if c.GetAIProvider() != "openai" {
		t.Fatalf("expected provider")
	}
	if c.GetAIModel() != "gpt-4" {
		t.Fatalf("expected model")
	}
	if c.GetOutputFormat() != "json" {
		t.Fatalf("expected format")
	}
}

func TestAppConfig_Empty(t *testing.T) {
	c := AppConfig{}
	if c.GetServerURL() != "" {
		t.Fatal("expected empty server URL")
	}
	if c.GetServerToken() != "" {
		t.Fatal("expected empty token")
	}
}

// --- TagDetail ---

func TestTagDetail_UnmarshalJSON(t *testing.T) {
	raw := `{
		"id": "t1",
		"slug_name": "go-lang",
		"display_name": "Go 语言",
		"question_count": 10,
		"follow_count": 5,
		"reserved": false,
		"recommend": true
	}`
	var td TagDetail
	if err := json.Unmarshal([]byte(raw), &td); err != nil {
		t.Fatal(err)
	}
	if td.SlugName != "go-lang" {
		t.Fatalf("expected go-lang, got %s", td.SlugName)
	}
	if td.DisplayName != "Go 语言" {
		t.Fatalf("expected Go 语言, got %s", td.DisplayName)
	}
	if td.QuestionCount != 10 {
		t.Fatalf("expected 10, got %d", td.QuestionCount)
	}
	if !td.Recommend {
		t.Fatal("expected recommend true")
	}
	if td.Reserved {
		t.Fatal("expected reserved false")
	}
}
