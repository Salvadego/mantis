package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type DashboardService struct {
	client *Client
}

func (s *DashboardService) GetReport(ctx context.Context) ([]TicketResponse, error) {

	endpoint := "/api/odata/cam/core/fh/v1/ReportAgingService"

	// TODO: Add from and to dates
	filter := fmt.Sprintf(
		"$filter=Filter_RSC eq true and Filter_Type_S_User eq 'DM' and Filter_S_User_ID eq 'S0011966293' and Filter_Change_At_From eq 2025-06-17T00:00:00Z and Filter_Change_At_To eq 2025-07-02T23:59:59Z",
	)

	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))
	headers := map[string]string{
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []TicketResponse `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
