package request

type FetchRequest struct {
	ApiKey       string `json:"api_key"`
	PartitionKey string `json:"partition_key"`
	RangeKey     string `json:"range_key"`
}
