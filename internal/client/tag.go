package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type TagService struct {
	client *Client
}

func NewTagService(client *Client) *TagService {
	return &TagService{client: client}
}

func (s *TagService) Page(ctx context.Context, req *models.TagListReq) (*models.TagListResp, error) {
	params := map[string]string{
		"page":      strconv.Itoa(req.Page),
		"page_size": strconv.Itoa(req.PageSize),
	}
	if req.QueryCond != "" {
		params["query_cond"] = req.QueryCond
	}
	var result models.TagListResp
	if err := s.client.GetJSON(ctx, "/tags/page", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *TagService) Get(ctx context.Context, slugName string) (*models.TagDetail, error) {
	var result models.TagDetail
	if err := s.client.GetJSON(ctx, "/tag", map[string]string{"slug_name": slugName}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *TagService) Search(ctx context.Context, prefix string) ([]*models.TagItem, error) {
	var result []*models.TagItem
	if err := s.client.GetJSON(ctx, "/question/tags", map[string]string{"tag": prefix}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TagService) Add(ctx context.Context, req *models.TagAddReq) (*models.TagDetail, error) {
	var result models.TagDetail
	if err := s.client.PostJSON(ctx, "/tag", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *TagService) Update(ctx context.Context, req *models.TagUpdateReq) error {
	return s.client.PutJSON(ctx, "/tag", req, nil)
}

func (s *TagService) Delete(ctx context.Context, slugName string) error {
	return s.client.DeleteJSON(ctx, "/tag", &models.TagDeleteReq{SlugName: slugName})
}
