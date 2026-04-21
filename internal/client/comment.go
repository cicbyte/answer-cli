package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type CommentService struct {
	client *Client
}

func NewCommentService(client *Client) *CommentService {
	return &CommentService{client: client}
}

func (s *CommentService) Page(ctx context.Context, req *models.CommentListReq) (*models.CommentListResp, error) {
	params := map[string]string{
		"object_id": req.ObjectID,
		"page":      strconv.Itoa(req.Page),
		"page_size": strconv.Itoa(req.PageSize),
	}
	var result models.CommentListResp
	if err := s.client.GetJSON(ctx, "/comment/page", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CommentService) Get(ctx context.Context, id string) (*models.CommentInfo, error) {
	var result models.CommentInfo
	if err := s.client.GetJSON(ctx, "/comment", map[string]string{"id": id}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CommentService) Add(ctx context.Context, req *models.CommentAddReq) (*models.CommentInfo, error) {
	var result models.CommentInfo
	if err := s.client.PostJSON(ctx, "/comment", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CommentService) Update(ctx context.Context, req *models.CommentUpdateReq) error {
	return s.client.PutJSON(ctx, "/comment", req, nil)
}

func (s *CommentService) Delete(ctx context.Context, id string) error {
	return s.client.DeleteJSON(ctx, "/comment", &models.CommentDeleteReq{ID: id})
}
