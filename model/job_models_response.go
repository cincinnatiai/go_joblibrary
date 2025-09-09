package model

type JobModelsResponse struct {
	Results          []*JobModel `json:"results"`
	LastPartitionKey *string     `json:"last_partition_key"`
	LastRangeKey     *string     `json:"last_range_key"`
}
