package client

import (
	"context"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/models"
)

type NotificationService struct {
	client *Client
}

func NewNotificationService(client *Client) *NotificationService {
	return &NotificationService{client: client}
}

func (s *NotificationService) Status(ctx context.Context) (*models.NotificationStatusResp, error) {
	var result models.NotificationStatusResp
	if err := s.client.GetJSON(ctx, "/notification/status", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *NotificationService) Page(ctx context.Context, req *models.NotificationListReq) (*models.NotificationListResp, error) {
	params := map[string]string{
		"page": strconv.Itoa(req.Page),
		"size": strconv.Itoa(req.Size),
	}
	if req.Type != "" {
		params["type"] = req.Type
	}
	if req.InboxType != "" {
		params["inbox_type"] = req.InboxType
	}
	var result models.NotificationListResp
	if err := s.client.GetJSON(ctx, "/notification/page", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *NotificationService) ClearAll(ctx context.Context, nType string) error {
	return s.client.PutJSON(ctx, "/notification/read/state/all", &models.NotificationReadAllReq{Type: nType}, nil)
}

func (s *NotificationService) ClearID(ctx context.Context, id string) error {
	return s.client.PutJSON(ctx, "/notification/read/state", &models.NotificationReadReq{ID: id}, nil)
}
