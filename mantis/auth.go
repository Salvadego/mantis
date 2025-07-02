package mantis

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type AuthService struct {
	client *Client
}

func (s *AuthService) Authenticate(ctx context.Context) error {
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("client_id", s.client.authConfig.ClientID)
	form.Add("client_secret", s.client.authConfig.ClientSecret)
	form.Add("username", s.client.authConfig.Username)
	form.Add("password", s.client.authConfig.Password)

	body := strings.NewReader(form.Encode())

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := s.client.doRequest(ctx,
		http.MethodPost,
		"/oauth2/v1/token",
		body,
		headers)

	if err != nil {
		return err
	}

	var tokenResp TokenResponse
	if err := parseResponse(resp, &tokenResp); err != nil {
		return err
	}

	s.client.SetToken(tokenResp.AccessToken)

	return nil
}
