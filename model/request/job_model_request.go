package request

import "github.com/cincinnatiai/go_joblibrary/model"

type JobModelCreateRequest struct {
	AccountId           string `json:"account_id"`
	UserId              string `json:"user_id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	Department          string `json:"department"`
	Category            string `json:"category"`
	Location            string `json:"location"`
	JobType             string `json:"job_type"`
	SalaryMin           int    `json:"salary_min"`
	SalaryMax           int    `json:"salary_max"`
	Requirements        string `json:"requirements"`
	Responsibilities    string `json:"responsibilities"`
	Benefits            string `json:"benefits"`
	PostedBy            string `json:"posted_by"`
	ApplicationDeadline string `json:"application_deadline"`
	ExperienceLevel     string `json:"experience_level"`
	RemoteAllowed       bool   `json:"remote_allowed"`
	ApiKey              string `json:"api_key"`
}

type JobModelFetchAllRequest struct {
	AccountId        string  `json:"account_id"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchByCategoryRequest struct {
	Category         string  `json:"category"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchByDepartmentRequest struct {
	Department       string  `json:"department"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchByStatusRequest struct {
	Status           string  `json:"status"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelUpdateRequest struct {
	Job    model.JobModel `json:"job"`
	UserId string         `json:"user_id"`
	ApiKey string         `json:"api_key"`
}

type JobModelFetchByLocationRequest struct {
	Location         string  `json:"location"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchByExperienceLevelRequest struct {
	ExperienceLevel  string  `json:"experience_level"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchByRemoteRequest struct {
	RemoteAllowed    bool    `json:"remote_allowed"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}

type JobModelFetchBySalaryRangeRequest struct {
	MinSalary        int     `json:"min_salary"`
	MaxSalary        int     `json:"max_salary"`
	AccountId        string  `json:"account_id,omitempty"`
	UserId           string  `json:"user_id,omitempty"`
	LastPartitionKey *string `json:"last_partition_key,omitempty"`
	LastRangeKey     *string `json:"last_range_key,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
	ApiKey           string  `json:"api_key"`
}
