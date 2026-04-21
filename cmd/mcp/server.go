package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func runMcpServer() error {
	cfg := common.GetAppConfig()
	if cfg.GetServerURL() == "" {
		return fmt.Errorf("server not configured, run: answer-cli config set server.base_url <url>")
	}
	if cfg.GetServerToken() == "" {
		return fmt.Errorf("not authenticated, run: answer-cli auth login")
	}

	cli := client.NewClient(&client.Config{
		BaseURL: cfg.GetServerURL(),
		Token:   cfg.GetServerToken(),
	})

	s := server.NewMCPServer(
		"answer-cli",
		"0.1.0",
		server.WithToolCapabilities(true),
	)

	// question_search
	s.AddTool(mcp.NewTool("question_search",
		mcp.WithDescription("Search questions by keyword, tag, or username"),
		mcp.WithString("keyword", mcp.Description("Search keyword")),
		mcp.WithString("tag", mcp.Description("Filter by tag slug")),
		mcp.WithString("order", mcp.Description("Sort order: newest|active|hot|score|unanswered")),
		mcp.WithNumber("page", mcp.Description("Page number")),
		mcp.WithNumber("limit", mcp.Description("Page size, max 100")),
	), makeHandler(cli, handleQuestionSearch))

	// question_get
	s.AddTool(mcp.NewTool("question_get",
		mcp.WithDescription("Get question detail by ID"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
	), makeHandler(cli, handleQuestionGet))

	// answer_list
	s.AddTool(mcp.NewTool("answer_list",
		mcp.WithDescription("List answers for a question"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
		mcp.WithNumber("page", mcp.Description("Page number")),
		mcp.WithNumber("limit", mcp.Description("Page size")),
	), makeHandler(cli, handleAnswerList))

	// tag_search
	s.AddTool(mcp.NewTool("tag_search",
		mcp.WithDescription("Search tags by prefix"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Tag name prefix")),
	), makeHandler(cli, handleTagSearch))

	// tag_get
	s.AddTool(mcp.NewTool("tag_get",
		mcp.WithDescription("Get tag detail by slug name"),
		mcp.WithString("slug_name", mcp.Required(), mcp.Description("Tag slug name")),
	), makeHandler(cli, handleTagGet))

	// user_search
	s.AddTool(mcp.NewTool("user_search",
		mcp.WithDescription("Search users by username"),
		mcp.WithString("username", mcp.Required(), mcp.Description("Username to search")),
	), makeHandler(cli, handleUserSearch))

	log.Info("MCP Server starting")

	return server.ServeStdio(s)
}

type handlerFunc func(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error)

func makeHandler(cli *client.Client, fn handlerFunc) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return fn(ctx, cli, req)
	}
}

func handleQuestionSearch(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyword := req.GetString("keyword", "")
	tag := req.GetString("tag", "")
	order := req.GetString("order", "newest")
	page := int(req.GetFloat("page", 1))
	limit := int(req.GetFloat("limit", 20))

	result, err := cli.Question.Page(ctx, &models.QuestionListReq{
		Page: page, PageSize: limit, Order: order, Tag: tag,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	type item struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		AnswerCount int    `json:"answer_count"`
		VoteCount   int    `json:"vote_count"`
		CreatedAt   string `json:"created_at"`
		Tags        []string `json:"tags,omitempty"`
	}

	var items []item
	for _, q := range result.List {
		tags := make([]string, len(q.Tags))
		for i, t := range q.Tags {
			tags[i] = t.SlugName
		}
		items = append(items, item{
			ID: q.ID, Title: q.Title, AnswerCount: q.AnswerCount,
			VoteCount: q.VoteCount,
			CreatedAt: models.FormatTimestamp(q.CreatedAt).Format("2006-01-02"),
			Tags: tags,
		})
	}

	if keyword != "" {
		// also try full-text search
		searchResult, err := cli.Search.Search(ctx, &models.SearchReq{
			Query: keyword, Page: page, Size: limit, Order: order,
		})
		if err == nil && len(searchResult.List) > 0 {
			for _, s := range searchResult.List {
				items = append(items, item{
					ID: s.ID, Title: s.Title, AnswerCount: s.AnswerCount,
					VoteCount: s.VoteCount,
					CreatedAt: models.FormatTimestamp(s.CreatedAt).Format("2006-01-02"),
				})
			}
		}
	}

	if items == nil {
		items = []item{}
	}
	data, _ := json.Marshal(items)
	return mcp.NewToolResultText(string(data)), nil
}

func handleQuestionGet(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("question_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	q, err := cli.Question.Get(ctx, id)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	tags := make([]string, len(q.Tags))
	for i, t := range q.Tags {
		tags[i] = t.SlugName
	}

	result := map[string]interface{}{
		"id":              q.ID,
		"title":           q.Title,
		"content":         q.Content,
		"answer_count":    q.AnswerCount,
		"vote_count":      q.VoteCount,
		"view_count":      q.ViewCount,
		"comment_count":   q.CommentCount,
		"status":          q.Status,
		"created_at":      models.FormatTimestamp(q.CreatedAt).Format("2006-01-02"),
		"tags":            tags,
		"accepted_answer": q.AcceptedAnswerID,
	}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

func handleAnswerList(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	questionID, err := req.RequireString("question_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	page := int(req.GetFloat("page", 1))
	limit := int(req.GetFloat("limit", 20))

	result, err := cli.Answer.Page(ctx, &models.AnswerListReq{
		QuestionID: questionID, Page: page, PageSize: limit,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	type item struct {
		ID        string `json:"id"`
		Content   string `json:"content"`
		VoteCount int    `json:"vote_count"`
		Accepted  int    `json:"accepted"`
		CreatedAt string `json:"created_at"`
		Username  string `json:"username"`
	}

	var items []item
	for _, a := range result.List {
		username := ""
		if a.UserInfo != nil {
			username = a.UserInfo.DisplayName
			if username == "" {
				username = a.UserInfo.Username
			}
		}
		items = append(items, item{
			ID: a.ID, Content: a.Content, VoteCount: a.VoteCount,
			Accepted: a.Accepted,
			CreatedAt: models.FormatTimestamp(a.CreatedAt).Format("2006-01-02"),
			Username: username,
		})
	}
	if items == nil {
		items = []item{}
	}
	data, _ := json.Marshal(items)
	return mcp.NewToolResultText(string(data)), nil
}

func handleTagSearch(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	tags, err := cli.Tag.Search(ctx, query)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if tags == nil {
		tags = []*models.TagItem{}
	}
	data, _ := json.Marshal(tags)
	return mcp.NewToolResultText(string(data)), nil
}

func handleTagGet(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	slug, err := req.RequireString("slug_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	t, err := cli.Tag.Get(ctx, slug)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := map[string]interface{}{
		"slug_name":      t.SlugName,
		"display_name":   t.DisplayName,
		"description":    t.OriginalText,
		"question_count": t.QuestionCount,
		"follow_count":   t.FollowCount,
	}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

func handleUserSearch(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	username, err := req.RequireString("username")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	users, err := cli.User.SearchUsers(ctx, username)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if users == nil {
		users = make([]*models.UserBasicInfo, 0)
	}
	data, _ := json.Marshal(users)
	return mcp.NewToolResultText(string(data)), nil
}
