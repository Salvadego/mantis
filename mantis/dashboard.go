package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type DashboardService struct {
	client *Client
}

type GetReportOptions struct {
	FilterRSC    bool
	FilterType   string
	FilterUserID string
	ChangeAtFrom *time.Time
	ChangeAtTo   *time.Time
}

func appendToList[Type comparable](list *[]Type, v Type) {
	*list = append(*list, v)
}

func (s *DashboardService) GetReport(
	ctx context.Context,
	opts *GetReportOptions,
) ([]TicketResponse, error) {
	endpoint := "/api/odata/cam/core/fh/v1/ReportAgingService"

	filterParts := []string{}
	if opts.FilterRSC {
		appendToList(
			&filterParts,
			"Filter_RSC eq true",
		)
	}

	if opts.FilterType != "" {
		appendToList(
			&filterParts,
			fmt.Sprintf("Filter_Type_S_User eq '%s'", opts.FilterType),
		)
	}

	if opts.FilterUserID != "" {
		appendToList(
			&filterParts,
			fmt.Sprintf("Filter_S_User_ID eq '%s'", opts.FilterUserID),
		)
	}

	if opts.ChangeAtFrom != nil {
		appendToList(
			&filterParts,
			fmt.Sprintf(
				"Filter_Change_At_From eq %s",
				opts.ChangeAtFrom.Format("2006-01-02T15:04:05Z"),
			),
		)
	}

	if opts.ChangeAtTo != nil {
		appendToList(
			&filterParts,
			fmt.Sprintf(
				"Filter_Change_At_To eq %s",
				opts.ChangeAtTo.Format("2006-01-02T15:04:05Z"),
			),
		)
	}

	var filter string
	if len(filterParts) > 0 {
		filter = "$filter=" + url.QueryEscape(joinFilterParts(filterParts, " and "))
	}

	path := fmt.Sprintf("%s?%s", endpoint, filter)
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

func joinFilterParts(parts []string, separator string) string {
	if len(parts) == 0 {
		return ""
	}

	if len(parts) == 1 {
		return parts[0]
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += separator + parts[i]
	}

	return result
}
