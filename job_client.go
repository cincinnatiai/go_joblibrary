package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_joblibrary/model"
	"go_joblibrary/model/request"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ClientConfig struct {
	BaseURL    string
	HTTPClient *http.Client
	ApiKey     string
}

type Client struct {
	config ClientConfig
}

func NewClient(config ClientConfig) *Client {
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}
	return &Client{config: config}
}

func NewClientWithDefaults(baseURL, apiKey string) *Client {
	return NewClient(ClientConfig{
		BaseURL:    baseURL,
		ApiKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	})
}

type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.config.BaseURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.config.HTTPClient.Do(req)
}

func (c *Client) makeGETRequest(endpoint string, params map[string]string) (*http.Response, error) {
	u, err := url.Parse(c.config.BaseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	for key, value := range params {
		if value != "" {
			q.Set(key, value)
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.config.HTTPClient.Do(req)
}

func processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if respErr := json.Unmarshal(body, &apiErr); respErr != nil {
			// If we can't unmarshal the error, create a generic one
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    string(body),
			}
		}
		apiErr.StatusCode = resp.StatusCode
		return &apiErr
	}

	if result != nil {
		if unmarshalErr := json.Unmarshal(body, result); unmarshalErr != nil {
			return fmt.Errorf("failed to unmarshal response: %w", unmarshalErr)
		}
	}

	return nil
}

// API Methods
func (c *Client) CreateJob(req request.JobModelCreateRequest) (*model.JobModel, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=create", req)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	var result model.JobModel
	if err := processResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FetchAllJobs(req request.JobModelFetchAllRequest) (*model.JobModelsResponse, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=fetchAll", req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all jobs: %w", err)
	}

	var result model.JobModelsResponse
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, processErr
	}

	return &result, nil
}

func (c *Client) FetchJobsByCategory(req request.JobModelFetchByCategoryRequest) (*model.JobModelsResponse, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=fetchByCategory", req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs by category: %w", err)
	}

	var result model.JobModelsResponse
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, processErr
	}

	return &result, nil
}

func (c *Client) FetchJobsByCategoryPublic(category string, lastPartitionKey, lastRangeKey *string) (*model.JobModelsResponse, error) {
	params := map[string]string{
		"category": category,
	}

	if lastPartitionKey != nil && *lastPartitionKey != "" {
		params["last_partition_key"] = *lastPartitionKey
	}
	if lastRangeKey != nil && *lastRangeKey != "" {
		params["last_range_key"] = *lastRangeKey
	}

	resp, err := c.makeGETRequest("?action=fetchByCategory", params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs by category (public): %w", err)
	}

	var result model.JobModelsResponse
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) FetchJobsByDepartment(req request.JobModelFetchByDepartmentRequest) (*model.JobModelsResponse, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=fetchByDepartment", req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jobs by department: %w", err)
	}

	var result model.JobModelsResponse
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, processErr
	}

	return &result, nil
}

func (c *Client) FetchJob(req request.FetchRequest) (*model.JobModel, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=fetch", req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job: %w", err)
	}

	var result model.JobModel
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, processErr
	}

	return &result, nil
}

func (c *Client) UpdateJob(req request.JobModelUpdateRequest) (*bool, error) {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=update", req)
	if err != nil {
		return nil, fmt.Errorf("failed to update job: %w", err)
	}

	var result bool
	if processErr := processResponse(resp, &result); processErr != nil {
		return nil, processErr
	}

	return &result, nil
}

func (c *Client) DeleteJob(req request.DeleteRequest) error {
	if req.ApiKey == "" {
		req.ApiKey = c.config.ApiKey
	}

	resp, err := c.makeRequest("POST", "?action=delete", req)
	if err != nil {
		return fmt.Errorf("failed to delete job: %w", err)
	}

	return processResponse(resp, nil)
}

func (c *Client) CreateJobSimple(accountId, userId, title, description, category, department string) (*model.JobModel, error) {
	req := request.JobModelCreateRequest{
		AccountId:   accountId,
		UserId:      userId,
		Title:       title,
		Description: description,
		Category:    category,
		Department:  department,
		ApiKey:      c.config.ApiKey,
	}

	return c.CreateJob(req)
}

func (c *Client) FetchJobSimple(partitionKey, rangeKey string) (*model.JobModel, error) {
	req := request.FetchRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		ApiKey:       c.config.ApiKey,
	}

	return c.FetchJob(req)
}

func (c *Client) DeleteJobSimple(partitionKey, rangeKey string) error {
	req := request.DeleteRequest{
		PartitionKey: partitionKey,
		RangeKey:     rangeKey,
		ApiKey:       c.config.ApiKey,
	}

	return c.DeleteJob(req)
}
