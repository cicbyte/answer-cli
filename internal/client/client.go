package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Client struct {
	client        *resty.Client
	baseURL       string
	token         string
	Comment       *CommentService
	Tag           *TagService
	Search        *SearchService
	Notification  *NotificationService
	User          *UserService
	Vote          *VoteService
	Question      *QuestionService
	Answer        *AnswerService
}

type Config struct {
	BaseURL string
	Token   string
	Timeout time.Duration
}

func NewClient(cfg *Config) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}

	client := resty.New().
		SetBaseURL(cfg.BaseURL + "/answer/api/v1").
		SetTimeout(cfg.Timeout).
		SetRetryCount(3).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(5 * time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if cfg.Token != "" {
		client.SetHeader("Authorization", cfg.Token)
	}

	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		log.Debug("API request",
			zap.String("method", req.Method),
			zap.String("url", req.URL),
		)
		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		log.Debug("API response",
			zap.Int("status", resp.StatusCode()),
			zap.String("url", resp.Request.URL),
			zap.Duration("duration", resp.Time()),
		)
		return nil
	})

	c := &Client{
		client:  client,
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
	}
	c.Comment = NewCommentService(c)
	c.Tag = NewTagService(c)
	c.Search = NewSearchService(c)
	c.Notification = NewNotificationService(c)
	c.User = NewUserService(c)
	c.Vote = NewVoteService(c)
	c.Question = NewQuestionService(c)
	c.Answer = NewAnswerService(c)
	return c
}

func (c *Client) SetToken(token string) {
	c.token = token
	c.client.SetHeader("Authorization", token)
}

func (c *Client) GetToken() string {
	return c.token
}

// RespBody Answer API 统一响应结构
type RespBody struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

// unwrap 解包 RespBody，将 data 字段反序列化到 result
func unwrap(resp *resty.Response, err error, result interface{}) error {
	if err != nil {
		log.Error("API request failed", zap.Error(err))
		return fmt.Errorf("request failed: %w", err)
	}

	var body RespBody
	if jsonErr := json.Unmarshal(resp.Body(), &body); jsonErr != nil {
		return fmt.Errorf("failed to parse response: %w", jsonErr)
	}

	if body.Code != 0 && body.Code != 200 {
		msg := body.Message
		if msg == "" {
			msg = body.Reason
		}
		return &APIError{Code: body.Code, Message: msg}
	}

	if result != nil && len(body.Data) > 0 {
		if jsonErr := json.Unmarshal(body.Data, result); jsonErr != nil {
			return fmt.Errorf("failed to parse response data: %w", jsonErr)
		}
	}

	return nil
}

// GetJSON GET 请求并解包 RespBody
func (c *Client) GetJSON(ctx context.Context, path string, params map[string]string, result interface{}) error {
	req := c.client.R().SetContext(ctx)
	if params != nil {
		req.SetQueryParams(params)
	}
	resp, err := req.Get(path)
	return unwrap(resp, err, result)
}

// PostJSON POST 请求并解包 RespBody
func (c *Client) PostJSON(ctx context.Context, path string, body interface{}, result interface{}) error {
	req := c.client.R().SetContext(ctx)
	if body != nil {
		req.SetBody(body)
	}
	resp, err := req.Post(path)
	return unwrap(resp, err, result)
}

// PutJSON PUT 请求并解包 RespBody
func (c *Client) PutJSON(ctx context.Context, path string, body interface{}, result interface{}) error {
	req := c.client.R().SetContext(ctx)
	if body != nil {
		req.SetBody(body)
	}
	resp, err := req.Put(path)
	return unwrap(resp, err, result)
}

// DeleteJSON DELETE 请求并解包 RespBody
func (c *Client) DeleteJSON(ctx context.Context, path string, body interface{}) error {
	req := c.client.R().SetContext(ctx)
	if body != nil {
		req.SetBody(body)
	}
	resp, err := req.Delete(path)
	return unwrap(resp, err, nil)
}

// APIError Answer API 错误
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error [%d]: %s", e.Code, e.Message)
}

func IsUnauthorizedError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == 401
	}
	return false
}

func IsNotFoundError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Code == 404
	}
	return false
}

func IsNetworkError(err error) bool {
	return err != nil && (err == http.ErrHandlerTimeout || err == context.DeadlineExceeded)
}
