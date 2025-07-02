package mantis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type TimesheetService struct {
	client *Client
	mu     sync.Mutex
}

func (s *TimesheetService) Create(
	ctx context.Context,
	timesheet Timesheet) (ErrorsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	reqBody, err := json.Marshal(timesheet)
	if err != nil {
		return ErrorsResponse{}, fmt.Errorf("falha ao serializar timesheet: %w", err)
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(
		ctx,
		http.MethodPost,
		"/api/odata/cam/core/system/v1/MTS_Timesheets",
		bytes.NewReader(reqBody),
		headers,
	)

	var errors ErrorsResponse
	err = json.NewDecoder(resp.Body).Decode(&errors)

	return errors, err
}

func (s *TimesheetService) GetTimesheets(
	ctx context.Context,
	userID int) ([]TimesheetsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()

	var startDate, endDate time.Time

	if currentDay > 26 {
		startDate = time.Date(currentYear, currentMonth, 26, 0, 0, 0, 0, time.UTC)
		endDate = startDate.AddDate(0, 1, -1)
	} else {
		endDate = time.Date(currentYear, currentMonth, 25, 0, 0, 0, 0, time.UTC)
		startDate = endDate.AddDate(0, -1, 1)
	}

	startDateStr := startDate.Format("2006-01-02T00:00:00Z")
	endDateStr := endDate.Format("2006-01-02T00:00:00Z")

	endpoint := "api/odata/cam/core/system/v1/MTS_Timesheets"
	filter := fmt.Sprintf(
		"$filter=userId eq %d and (dateDoc ge %s and dateDoc le %s)",
		userID,
		startDateStr,
		endDateStr)

	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	headers := map[string]string{
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []TimesheetsResponse `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

func (s *TimesheetService) GetProjectTimesheets(
	ctx context.Context,
	employeeId int) ([]ProjectTimesheet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	headers := map[string]string{
		"SourceSystem": "APP",
	}

	endpoint := "/api/odata/cam/core/system/v1/MTS_ProjectTimesheet"

	filter := fmt.Sprintf("$filter=employeeId eq %d", employeeId)
	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}
	var result struct {
		Value []ProjectTimesheet `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}

func (s *TimesheetService) GetProjectTimesheetsWithTicketAllocation(
	ctx context.Context,
	empID int) ([]ProjectTimesheet, error) {
	result, err := s.GetProjectTimesheets(ctx, empID)
	if err != nil {
		return nil, err
	}

	filtered := filter(result, func(project ProjectTimesheet) bool {
		return project.ProjectNeedTicket
	})

	if len(filtered) == 0 {
		return nil, fmt.Errorf(
			"projetos encontrados, mas faltam alocações de tíquetes em alguns projetos")
	}

	return filtered, nil
}

func (s *TimesheetService) DeleteTimesheet(
	ctx context.Context,
	timesheetID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	endpoint := fmt.Sprintf(
		"/api/odata/cam/core/system/v1/MTS_Timesheets(%d)",
		timesheetID)

	headers := map[string]string{
		"Content-Type": "application/json",
		"SourceSystem": "APP",
	}

	resp, err := s.client.doRequest(
		ctx,
		http.MethodDelete,
		endpoint,
		nil,
		headers,
	)
	if err != nil {
		return err
	}

	return parseResponse(resp, nil)
}
