package client

import (
	"context"

	"github.com/cicbyte/answer-cli/internal/models"
)

type VoteService struct {
	client *Client
}

func NewVoteService(client *Client) *VoteService {
	return &VoteService{client: client}
}

func (s *VoteService) VoteUp(ctx context.Context, objectID string) error {
	return s.client.PostJSON(ctx, "/vote/up", &models.VoteReq{ObjectID: objectID}, nil)
}

func (s *VoteService) VoteDown(ctx context.Context, objectID string) error {
	return s.client.PostJSON(ctx, "/vote/down", &models.VoteReq{ObjectID: objectID}, nil)
}
