package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

type TicketService struct {
	client *Client
	mu     sync.Mutex
}

func (s *TicketService) GetTickets(
	ctx context.Context,
	userID int,
) ([]TicketsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	endpoint := "api/odata/cam/core/system/v1/MTS_SMTickets"
	filter := fmt.Sprintf(
		"$filter=userId eq %d",
		userID)

	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	headers := map[string]string{
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []TicketsResponse `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
