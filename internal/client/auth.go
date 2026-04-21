package client

import (
	"context"

	"github.com/cicbyte/answer-cli/internal/models"
)

type AuthService struct {
	client *Client
}

func NewAuthService(client *Client) *AuthService {
	return &AuthService{client: client}
}

func (s *AuthService) Login(ctx context.Context, email, pass string) (*models.LoginResponse, error) {
	req := map[string]string{
		"e_mail": email,
		"pass":   pass,
	}
	var result models.LoginResponse
	if err := s.client.PostJSON(ctx, "/user/login/email", req, &result); err != nil {
		return nil, err
	}
	if result.AccessToken != "" {
		s.client.SetToken(result.AccessToken)
	}
	return &result, nil
}

func (s *AuthService) Logout(ctx context.Context) error {
	if err := s.client.GetJSON(ctx, "/user/logout", nil, nil); err != nil {
		return err
	}
	s.client.SetToken("")
	return nil
}

func (s *AuthService) GetCurrentUser(ctx context.Context) (*models.UserInfoResp, error) {
	var result models.UserInfoResp
	if err := s.client.GetJSON(ctx, "/user/info", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AuthService) IsAuthenticated() bool {
	return s.client.GetToken() != ""
}
