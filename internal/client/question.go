package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type QuestionService struct {
	client *Client
}

func NewQuestionService(client *Client) *QuestionService {
	return &QuestionService{client: client}
}

func (s *QuestionService) Page(ctx context.Context, req *models.QuestionListReq) (*models.QuestionPageResp, error) {
	params := map[string]string{
		"page":      strconv.Itoa(req.Page),
		"page_size": strconv.Itoa(req.PageSize),
	}
	if req.Order != "" {
		params["order"] = req.Order
	}
	if req.Tag != "" {
		params["tag"] = req.Tag
	}
	if req.Username != "" {
		params["username"] = req.Username
	}
	if req.InDays > 0 {
		params["in_days"] = strconv.Itoa(req.InDays)
	}
	var result models.QuestionPageResp
	if err := s.client.GetJSON(ctx, "/question/page", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *QuestionService) Get(ctx context.Context, id string) (*models.QuestionInfoResp, error) {
	var result models.QuestionInfoResp
	if err := s.client.GetJSON(ctx, "/question/info", map[string]string{"id": id}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *QuestionService) Add(ctx context.Context, req *models.QuestionAddReq) (*models.QuestionInfoResp, error) {
	var result models.QuestionInfoResp
	if err := s.client.PostJSON(ctx, "/question", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *QuestionService) Update(ctx context.Context, req *models.QuestionUpdateReq) error {
	return s.client.PutJSON(ctx, "/question", req, nil)
}

func (s *QuestionService) Delete(ctx context.Context, id string) error {
	return s.client.DeleteJSON(ctx, "/question", &models.QuestionDeleteReq{ID: id})
}

func (s *QuestionService) Close(ctx context.Context, id string, closeType int, closeMsg string) error {
	req := &models.QuestionCloseReq{ID: id, CloseType: closeType, CloseMsg: closeMsg}
	return s.client.PutJSON(ctx, "/question/status", req, nil)
}

func (s *QuestionService) Reopen(ctx context.Context, id string) error {
	return s.client.PutJSON(ctx, "/question/reopen", &models.QuestionReopenReq{ID: id}, nil)
}
