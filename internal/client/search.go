package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type SearchService struct {
	client *Client
}

func NewSearchService(client *Client) *SearchService {
	return &SearchService{client: client}
}

func (s *SearchService) Search(ctx context.Context, req *models.SearchReq) (*models.SearchResp, error) {
	params := map[string]string{
		"q":    req.Query,
		"page": strconv.Itoa(req.Page),
		"size": strconv.Itoa(req.Size),
	}
	if req.Order != "" {
		params["order"] = req.Order
	}
	var result models.SearchResp
	if err := s.client.GetJSON(ctx, "/search", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
