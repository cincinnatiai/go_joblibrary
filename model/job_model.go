package model

type JobModel struct {
	// AccountPK_AccountRK
	PartitionKey string `json:"partition_key"`
	// Timestamp so we can sort
	RangeKey            string `json:"range_key"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	Category            string `json:"category"` // Should be GSI
	Department          string `json:"department"`
	Location            string `json:"location"`
	JobType             string `json:"job_type"` // full-time, part-time, contract, etc.
	SalaryMin           int    `json:"salary_min"`
	SalaryMax           int    `json:"salary_max"`
	Requirements        string `json:"requirements"`
	Responsibilities    string `json:"responsibilities"`
	Benefits            string `json:"benefits"`
	Status              string `json:"status"`    // active, inactive, filled, etc.
	PostedBy            string `json:"posted_by"` // user who posted the job
	ApplicationDeadline string `json:"application_deadline"`
	ExperienceLevel     string `json:"experience_level"` // entry, mid, senior, etc.
	RemoteAllowed       bool   `json:"remote_allowed"`
	ApplicationCount    int    `json:"application_count"`
	ViewCount           int    `json:"view_count"`
	Created             string `json:"created"`
	Modified            string `json:"modified"`
}
