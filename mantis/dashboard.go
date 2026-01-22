package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DashboardService struct {
	client *Client
}

type GetReportOptions struct {
	FilterRSC        bool
	FilterType       string
	FilterUserID     string
	FilterContractID string
	ChangeAtFrom     *time.Time
	ChangeAtTo       *time.Time
}

func appendToList[Type comparable](list *[]Type, v Type) {
	*list = append(*list, v)
}

func defaultGetReportOptions() *GetReportOptions {
	now := time.Now().UTC()
	from := now.AddDate(0, 0, -15)

	return &GetReportOptions{
		FilterRSC:        true,
		FilterType:       "",
		FilterUserID:     "",
		FilterContractID: "",
		ChangeAtFrom:     &from,
		ChangeAtTo:       &now,
	}
}

func mergeGetReportOptions(opts *GetReportOptions) *GetReportOptions {
	def := defaultGetReportOptions()
	if opts == nil {
		return def
	}

	if opts.FilterRSC {
		def.FilterRSC = true
	}
	if opts.FilterType != "" {
		def.FilterType = opts.FilterType
	}
	if opts.FilterContractID != "" {
		def.FilterContractID = opts.FilterContractID
	}
	if opts.FilterUserID != "" {
		def.FilterUserID = opts.FilterUserID
	}
	if opts.ChangeAtFrom != nil {
		def.ChangeAtFrom = opts.ChangeAtFrom
	}
	if opts.ChangeAtTo != nil {
		def.ChangeAtTo = opts.ChangeAtTo
	}

	return def
}

func (s *DashboardService) GetReport(
	ctx context.Context,
	opts *GetReportOptions,
) ([]TicketResponse, error) {

	opts = mergeGetReportOptions(opts)
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

	if opts.FilterContractID != "" {
		appendToList(
			&filterParts,
			fmt.Sprintf("Filter_Contract_ID eq '%s'", opts.FilterContractID),
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
		filter = "$filter=" + url.PathEscape(strings.Join(filterParts, " and "))
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

func (s *DashboardService) GetReportContracts(
	ctx context.Context,
) ([]LtContract, error) {

	endpoint := "/api/odata/cam/core/fh/v1/ReportContracts"

	headers := map[string]string{
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(ctx, http.MethodGet, endpoint, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []ContractResponse `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value[0].LtContracts, nil
}
