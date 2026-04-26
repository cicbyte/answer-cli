package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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

	s.AddTool(mcp.NewTool("question_search",
		mcp.WithDescription("Search questions by keyword, tag, or sort order"),
		mcp.WithString("keyword", mcp.Description("Search keyword for full-text search")),
		mcp.WithString("tag", mcp.Description("Filter by tag slug")),
		mcp.WithString("order", mcp.Description("Sort: newest|active|hot|score|unanswered|relevance")),
		mcp.WithNumber("page", mcp.Description("Page number")),
		mcp.WithNumber("limit", mcp.Description("Page size, max 100")),
	), makeHandler(cli, handleQuestionSearch))

	s.AddTool(mcp.NewTool("question_get",
		mcp.WithDescription("Get question detail by ID, including content, answers and comments"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
	), makeHandler(cli, handleQuestionGet))

	s.AddTool(mcp.NewTool("question_create",
		mcp.WithDescription("Create a new question"),
		mcp.WithString("title", mcp.Required(), mcp.Description("Question title")),
		mcp.WithString("content", mcp.Required(), mcp.Description("Question content in Markdown")),
		mcp.WithString("tags", mcp.Description("Comma-separated tag slugs, e.g. \"go,linux\"")),
	), makeHandler(cli, handleQuestionCreate))

	s.AddTool(mcp.NewTool("question_update",
		mcp.WithDescription("Update an existing question's title, content or tags"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
		mcp.WithString("title", mcp.Description("New title")),
		mcp.WithString("content", mcp.Description("New content in Markdown")),
		mcp.WithString("tags", mcp.Description("Comma-separated tag slugs")),
	), makeHandler(cli, handleQuestionUpdate))

	s.AddTool(mcp.NewTool("answer_list",
		mcp.WithDescription("List answers for a specific question"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
		mcp.WithNumber("limit", mcp.Description("Max results, default 20")),
	), makeHandler(cli, handleAnswerList))

	s.AddTool(mcp.NewTool("answer_create",
		mcp.WithDescription("Create an answer for a question"),
		mcp.WithString("question_id", mcp.Required(), mcp.Description("Question ID")),
		mcp.WithString("content", mcp.Required(), mcp.Description("Answer content in Markdown")),
	), makeHandler(cli, handleAnswerCreate))

	s.AddTool(mcp.NewTool("answer_update",
		mcp.WithDescription("Update an existing answer"),
		mcp.WithString("answer_id", mcp.Required(), mcp.Description("Answer ID")),
		mcp.WithString("content", mcp.Required(), mcp.Description("New content in Markdown")),
	), makeHandler(cli, handleAnswerUpdate))

	s.AddTool(mcp.NewTool("comment_add",
		mcp.WithDescription("Add a comment to a question or answer"),
		mcp.WithString("object_id", mcp.Required(), mcp.Description("Question or Answer ID")),
		mcp.WithString("content", mcp.Required(), mcp.Description("Comment content")),
	), makeHandler(cli, handleCommentAdd))

	s.AddTool(mcp.NewTool("tag_search",
		mcp.WithDescription("Search tags by prefix"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Tag name prefix")),
	), makeHandler(cli, handleTagSearch))

	log.Info("MCP Server starting")

	return server.ServeStdio(s)
}

type handlerFunc func(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error)

func makeHandler(cli *client.Client, fn handlerFunc) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return fn(ctx, cli, req)
	}
}

// ── Question ──

func handleQuestionSearch(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyword := req.GetString("keyword", "")
	tag := req.GetString("tag", "")
	order := req.GetString("order", "newest")
	page := int(req.GetFloat("page", 1))
	limit := int(req.GetFloat("limit", 20))
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	type item struct {
		ID          string   `json:"id"`
		Title       string   `json:"title"`
		AnswerCount int      `json:"answer_count"`
		VoteCount   int      `json:"vote_count"`
		CreatedAt   string   `json:"created_at"`
		Tags        []string `json:"tags,omitempty"`
	}

	var items []item
	if keyword != "" {
		if order == "" || order == "newest" {
			order = "relevance"
		}
		result, err := cli.Search.Search(ctx, &models.SearchReq{
			Query: keyword, Page: page, Size: limit, Order: order,
		})
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		for _, s := range result.List {
			if s.Object == nil {
				continue
			}
			items = append(items, item{
				ID: s.Object.ID, Title: s.Object.Title,
				AnswerCount: s.Object.AnswerCount, VoteCount: s.Object.VoteCount,
				CreatedAt: models.FormatTimestamp(s.Object.CreatedAt).Format("2006-01-02"),
			})
		}
	}

	listResult, err := cli.Question.Page(ctx, &models.QuestionListReq{
		Page: page, PageSize: limit, Order: order, Tag: tag,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	for _, q := range listResult.List {
		tags := make([]string, len(q.Tags))
		for i, t := range q.Tags {
			tags[i] = t.SlugName
		}
		items = append(items, item{
			ID: q.ID, Title: q.Title, AnswerCount: q.AnswerCount,
			VoteCount: q.VoteCount,
			CreatedAt: models.FormatTimestamp(q.CreatedAt).Format("2006-01-02"),
			Tags:      tags,
		})
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

	result := map[string]any{
		"id": q.ID, "title": q.Title, "content": q.Content,
		"answer_count": q.AnswerCount, "vote_count": q.VoteCount,
		"view_count": q.ViewCount, "comment_count": q.CommentCount,
		"created_at": models.FormatTimestamp(q.CreatedAt).Format("2006-01-02"),
		"tags": tags, "accepted_answer": q.AcceptedAnswerID,
	}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

func handleQuestionCreate(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	title, err := req.RequireString("title")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	content, err := req.RequireString("content")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var tags []models.TagAddReq
	if tagStr := req.GetString("tags", ""); tagStr != "" {
		for _, slug := range strings.Split(tagStr, ",") {
			slug = strings.TrimSpace(slug)
			if slug != "" {
				tags = append(tags, models.TagAddReq{SlugName: slug, DisplayName: slug})
			}
		}
	}

	q, err := cli.Question.Add(ctx, &models.QuestionAddReq{
		Title: title, Content: content, Tags: tags,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := map[string]any{"id": q.ID, "title": q.Title}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

func handleQuestionUpdate(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("question_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	updateReq := models.QuestionUpdateReq{ID: id}
	if title := req.GetString("title", ""); title != "" {
		updateReq.Title = title
	}
	if content := req.GetString("content", ""); content != "" {
		updateReq.Content = content
	}
	if tagStr := req.GetString("tags", ""); tagStr != "" {
		updateReq.Tags = strings.Split(tagStr, ",")
		for i := range updateReq.Tags {
			updateReq.Tags[i] = strings.TrimSpace(updateReq.Tags[i])
		}
	}

	if err := cli.Question.Update(ctx, &updateReq); err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText("ok"), nil
}

// ── Answer ──

func handleAnswerList(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	questionID, err := req.RequireString("question_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	limit := int(req.GetFloat("limit", 20))

	result, err := cli.Answer.Page(ctx, &models.AnswerListReq{
		QuestionID: questionID, Page: 1, PageSize: limit,
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

func handleAnswerCreate(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	questionID, err := req.RequireString("question_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	content, err := req.RequireString("content")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	a, err := cli.Answer.Add(ctx, &models.AnswerAddReq{
		QuestionID: questionID, Content: content,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := map[string]any{"id": a.ID, "question_id": questionID}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

func handleAnswerUpdate(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	answerID, err := req.RequireString("answer_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	content, err := req.RequireString("content")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := cli.Answer.Update(ctx, &models.AnswerUpdateReq{
		ID: answerID, Content: content,
	}); err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText("ok"), nil
}

// ── Comment ──

func handleCommentAdd(ctx context.Context, cli *client.Client, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	objectID, err := req.RequireString("object_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	content, err := req.RequireString("content")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	c, err := cli.Comment.Add(ctx, &models.CommentAddReq{
		ObjectID: objectID, OriginalText: content,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := map[string]any{"comment_id": c.CommentID, "object_id": objectID}
	data, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(data)), nil
}

// ── Tag ──

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
