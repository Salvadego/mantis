package mantis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type EmployeeService struct {
	client *Client
	mu     sync.Mutex
}

func (s *EmployeeService) GetEmployeeById(
	ctx context.Context,
	userId int) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	headers := map[string]string{
		"SourceSystem": "APP",
	}
	endpoint := "/api/odata/cam/core/system/v1/MTS_Employees"

	filter := fmt.Sprintf("$top=50&$filter=userId eq %d and isActive eq true", userId)
	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return Employee{}, err
	}

	var result struct {
		Value []Employee `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return Employee{}, err
	}

	if len(result.Value) == 0 {
		return Employee{}, fmt.Errorf(
			"não foi possível encontrar funcionário ativo: %d",
			userId)
	}

	return result.Value[0], nil
}

func (s *EmployeeService) GetEmployeeByName(
	ctx context.Context,
	processorName string) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	headers := map[string]string{
		"SourceSystem": "APP",
	}
	endpoint := "/api/odata/cam/core/system/v1/MTS_Employees"

	nameParts := strings.Fields(strings.ToLower(processorName))

	var filterConditions []string

	for _, part := range nameParts {
		filterConditions = append(
			filterConditions,
			fmt.Sprintf("contains(tolower(fullName),'%s')", part))
	}

	filterExpr := strings.Join(filterConditions, " and ")
	filter := fmt.Sprintf("$top=50&$filter=(%s) and isActive eq true", filterExpr)
	path := fmt.Sprintf("%s?%s", endpoint, url.PathEscape(filter))

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return Employee{}, err
	}

	var result struct {
		Value []Employee `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return Employee{}, err
	}

	if len(result.Value) == 0 {
		return Employee{}, fmt.Errorf(
			"não foi possível encontrar funcionário ativo: %s",
			processorName)
	}

	return result.Value[0], nil
}

func (s *EmployeeService) GetEmployeeList(
	ctx context.Context,
	supervisorID *int,
) ([]S_Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	headers := map[string]string{
		"SourceSystem": "APP",
	}

	endpoint := "/api/odata/cam/core/fh/v1/EmployeeList"

	var queryParts []string

	if supervisorID != nil {
		queryParts = append(
			queryParts,
			fmt.Sprintf("$filter=Filter_Supervisor_ID eq %d", *supervisorID),
		)
	}

	query := strings.Join(queryParts, "&")
	path := endpoint
	if query != "" {
		path = fmt.Sprintf("%s?%s", endpoint, query)
	}

	resp, err := s.client.doRequest(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		return nil, err
	}

	var result struct {
		Value []S_Employee `json:"value"`
	}

	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
