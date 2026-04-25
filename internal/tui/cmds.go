package tui

import (
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/models"
)

const apiTimeout = 15 * time.Second

func apiCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), apiTimeout)
}

func safeCmd(f func() tea.Msg) tea.Cmd {
	return func() tea.Msg {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[tui panic recovered] %v\n", r)
			}
		}()
		return f()
	}
}

func loadQuestionsCmd(cli *client.Client, page, pageSize int, order string) tea.Cmd {
	return safeCmd(func() tea.Msg {
		ctx, cancel := apiCtx()
		defer cancel()
		resp, err := cli.Question.Page(ctx, &models.QuestionListReq{
			Page: page, PageSize: pageSize, Order: order,
		})
		if err != nil {
			return apiErrorMsg{err: err}
		}
		return questionsLoadedMsg{questions: resp.List, count: resp.Count, page: page}
	})
}

func loadQuestionDetailCmd(cli *client.Client, id string) tea.Cmd {
	return safeCmd(func() tea.Msg {
		ctx, cancel := apiCtx()
		defer cancel()
		q, err := cli.Question.Get(ctx, id)
		if err != nil {
			return apiErrorMsg{err: err}
		}
		return questionDetailLoadedMsg{question: q}
	})
}

func loadAnswersCmd(cli *client.Client, questionID string) tea.Cmd {
	return safeCmd(func() tea.Msg {
		ctx, cancel := apiCtx()
		defer cancel()
		resp, err := cli.Answer.Page(ctx, &models.AnswerListReq{
			QuestionID: questionID, Page: 1, PageSize: 50,
		})
		if err != nil {
			return apiErrorMsg{err: err}
		}
		return answersLoadedMsg{answers: resp.List, count: resp.Count}
	})
}

func loadCommentsCmd(cli *client.Client, objectID string) tea.Cmd {
	return safeCmd(func() tea.Msg {
		ctx, cancel := apiCtx()
		defer cancel()
		resp, err := cli.Comment.Page(ctx, &models.CommentListReq{
			ObjectID: objectID, Page: 1, PageSize: 50,
		})
		if err != nil {
			return apiErrorMsg{err: err}
		}
		return commentsLoadedMsg{objectID: objectID, comments: resp.List}
	})
}

func searchCmd(cli *client.Client, query string) tea.Cmd {
	return safeCmd(func() tea.Msg {
		ctx, cancel := apiCtx()
		defer cancel()
		resp, err := cli.Search.Search(ctx, &models.SearchReq{
			Query: query, Page: 1, Size: 20, Order: "relevance",
		})
		if err != nil {
			return apiErrorMsg{err: err}
		}
		return searchResultsMsg{results: resp.List, count: resp.Count, query: query}
	})
}

func checkAuthCmd(cli *client.Client) tea.Cmd {
	return safeCmd(func() tea.Msg {
		if cli.GetToken() == "" {
			return authCheckedMsg{username: ""}
		}
		ctx, cancel := apiCtx()
		defer cancel()
		auth := client.NewAuthService(cli)
		user, err := auth.GetCurrentUser(ctx)
		if err != nil {
			return authCheckedMsg{username: ""}
		}
		return authCheckedMsg{username: user.DisplayName}
	})
}

func NewTUIClient() (*client.Client, error) {
	cfg := common.GetAppConfig()
	if cfg.GetServerURL() == "" {
		return nil, fmt.Errorf("server not configured, run: answer-cli config set server.base_url <url>")
	}
	return client.NewClient(&client.Config{
		BaseURL: cfg.GetServerURL(),
		Token:   cfg.GetServerToken(),
	}), nil
}
