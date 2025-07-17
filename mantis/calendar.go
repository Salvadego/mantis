package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type CalendarService struct {
	client *Client
}

func (s *CalendarService) GetCalendar(
	ctx context.Context,
	year int,
	month time.Month,
) ([]NonBusinessDay, error) {

	endpoint := "/api/odata/cam/core/system/v1/MTS_NonBusinessDays"
	start := time.Date(
		year,
		month-1,
		26,
		0,
		0,
		0,
		0,
		time.UTC).Format("2006-01-02T00:00:00Z")

	end := time.Date(
		year,
		month,
		25,
		0,
		0,
		0,
		0,
		time.UTC).Format("2006-01-02T00:00:00Z")

	filter := fmt.Sprintf(
		"$filter=date ge %s and date le %s and (organizationId eq 0 or organizationId eq 1000073)",
		start,
		end,
	)

	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Value []NonBusinessDay `json:"value"`
	}
	if err := parseResponse(resp, &response); err != nil {
		return nil, err
	}

	return response.Value, nil
}
