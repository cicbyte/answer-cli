package client

import (
	"context"

	"github.com/cicbyte/answer-cli/internal/models"
)

// UserService 用户服务
type UserService struct {
	client *Client
}

// NewUserService 创建用户服务
func NewUserService(client *Client) *UserService {
	return &UserService{client: client}
}

// SearchUsers 搜索用户
func (s *UserService) SearchUsers(ctx context.Context, query string) ([]*models.UserBasicInfo, error) {
	var result []*models.UserBasicInfo
	if err := s.client.GetJSON(ctx, "/users/search", map[string]string{"query": query}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetByUsername 获取当前登录用户信息
func (s *UserService) GetByUsername(ctx context.Context) (*models.UserBasicInfo, error) {
	var result models.UserBasicInfo
	if err := s.client.GetJSON(ctx, "/user/info", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
