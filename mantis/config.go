package mantis

import "net/http"

type AuthConfig struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
}

type ClientConfig struct {
	BaseURL    string
	Language   string
	UserAgent  string
	HTTPClient *http.Client
}
