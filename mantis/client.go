package mantis

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL = "https://mantis-br.nttdata-solutions.com"
	defaultTimeout = 60 * time.Second
)

type Client struct {
	baseURL    string
	language   string
	roleID     string
	userAgent  string
	httpClient *http.Client
	authConfig AuthConfig

	tokenMu sync.RWMutex
	token   string

	Auth      *AuthService
	Calendar  *CalendarService
	Timesheet *TimesheetService
	Tickets   *TicketService
	Employee  *EmployeeService
	Dashboard *DashboardService
	Reference *ReferenceService
}

func NewClient(authConfig AuthConfig, clientConfig *ClientConfig) *Client {
	if clientConfig == nil {
		clientConfig = &ClientConfig{}
	}

	baseURL := clientConfig.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	httpClient := clientConfig.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
			Timeout: defaultTimeout,
		}
	}

	language := clientConfig.Language
	if language == "" {
		language = "pt_BR"
	}

	userAgent := clientConfig.UserAgent
	if userAgent == "" {
		userAgent = "mantis-go-client/1.0"
	}

	client := &Client{
		baseURL:    baseURL,
		language:   language,
		userAgent:  userAgent,
		httpClient: httpClient,
		authConfig: authConfig,
	}

	client.Auth = &AuthService{client: client}
	client.Timesheet = &TimesheetService{client: client}
	client.Employee = &EmployeeService{client: client}
	client.Dashboard = &DashboardService{client: client}
	client.Calendar = &CalendarService{client: client}
	client.Reference = &ReferenceService{client: client}
	client.Tickets = &TicketService{client: client}

	return client
}

func (c *Client) Token() string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.token
}

func (c *Client) SetToken(token string) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	c.token = token
}

func (c *Client) SetRoleID(roleID string) {
	c.roleID = roleID
}

func (c *Client) GetRoleID() string {
	return c.roleID
}

func (c *Client) doRequest(
	ctx context.Context,
	method,
	path string,
	body io.Reader,
	headers map[string]string) (*http.Response, error) {
	pathParts := strings.SplitN(path, "?", 2)
	basePath := pathParts[0]

	baseURL, err := url.JoinPath(c.baseURL, basePath)
	if err != nil {
		return nil, fmt.Errorf("URL inválida: %w", err)
	}

	reqURL := baseURL
	if len(pathParts) > 1 {
		reqURL = baseURL + "?" + pathParts[1]
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar requisição: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Language", c.language)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")

	token := c.Token()
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	if c.roleID != "" {
		req.Header.Set("RoleID", c.roleID)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNetworkError, err)
	}

	return resp, nil
}

func (c *Client) GetUserRoles(ctx context.Context, userID int) ([]UserRole, error) {
	headers := map[string]string{
		"SourceSystem": "APP",
	}

	endpoint := fmt.Sprintf(
		"/api/odata/cam/core/system/v1/UserBasicProfiles(%d)",
		userID,
	)

	resp, err := c.doRequest(ctx, http.MethodGet, endpoint, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value struct {
			Clients []struct {
				UserRoles []UserRole `json:"User_Roles"`
			} `json:"Clients"`
		} `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return []UserRole{}, err
	}

	if len(result.Value.Clients) == 0 {
		log.Println(result.Value)
		return []UserRole{}, fmt.Errorf("Empty clients")
	}

	return result.Value.Clients[0].UserRoles, nil
}

func parseResponse(resp *http.Response, successData any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	// log.Printf("response body: %s", body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var errResp ErrorsResponse
	if err := json.Unmarshal(body, &errResp); err == nil && errResp.IsError() {
		return &APIError{
			StatusCode: resp.StatusCode,
			Response:   errResp,
		}
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("http %s", resp.Status)
	}

	if successData != nil {
		if err := json.Unmarshal(body, successData); err != nil {
			return fmt.Errorf("failed to parse success body: %w", err)
		}
	}

	return nil
}
