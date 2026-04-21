package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type AnswerService struct {
	client *Client
}

func NewAnswerService(client *Client) *AnswerService {
	return &AnswerService{client: client}
}

func (s *AnswerService) Page(ctx context.Context, req *models.AnswerListReq) (*models.AnswerListResp, error) {
	params := map[string]string{
		"question_id": req.QuestionID,
		"page":        strconv.Itoa(req.Page),
		"page_size":   strconv.Itoa(req.PageSize),
	}
	if req.Order != "" {
		params["order"] = req.Order
	}
	var result models.AnswerListResp
	if err := s.client.GetJSON(ctx, "/answer/page", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AnswerService) Get(ctx context.Context, id string) (*models.AnswerInfo, error) {
	var result models.AnswerInfo
	if err := s.client.GetJSON(ctx, "/answer/info", map[string]string{"id": id}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AnswerService) Add(ctx context.Context, req *models.AnswerAddReq) (*models.AnswerInfo, error) {
	var result models.AnswerInfo
	if err := s.client.PostJSON(ctx, "/answer", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AnswerService) Update(ctx context.Context, req *models.AnswerUpdateReq) error {
	return s.client.PutJSON(ctx, "/answer", req, nil)
}

func (s *AnswerService) Delete(ctx context.Context, id string) error {
	return s.client.DeleteJSON(ctx, "/answer", &models.AnswerDeleteReq{ID: id})
}

func (s *AnswerService) Accept(ctx context.Context, questionID, answerID string) error {
	req := &models.AnswerAcceptReq{QuestionID: questionID, AnswerID: answerID}
	return s.client.PostJSON(ctx, "/answer/acceptance", req, nil)
}
