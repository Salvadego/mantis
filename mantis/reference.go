package mantis

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type ReferenceService struct {
	client *Client
}

type ReferenceTypeFilter struct {
	ColumnName string `json:"ColumnName"`
	TableName  string `json:"TableName"`
}

func (s *ReferenceService) GetReferenceTypes(ctx context.Context, filter ReferenceTypeFilter) ([]ReferenceType, error) {
	headers := map[string]string{
		"SourceSystem": "APP",
	}

	endpoint := "/api/odata/cam/core/fh/v1/ReferenceList"
	filterString := fmt.Sprintf("ColumnName eq '%s' and TableName eq '%s'", filter.ColumnName, filter.TableName)

	path := fmt.Sprintf("%s?$filter=%s", endpoint, filterString)

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []ReferenceType `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
