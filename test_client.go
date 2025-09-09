package main

import (
	"encoding/json"
	"fmt"
	"go_joblibrary/model"
	"go_joblibrary/model/request"
	"log"
	"os"
	"time"
)

func main() {
	config := struct {
		BaseURL   string
		ApiKey    string
		AccountId string
		UserId    string
	}{
		BaseURL:   "",
		ApiKey:    "",
		AccountId: "",
		UserId:    "",
	}

	// Add your environment variables before running
	if envURL := os.Getenv("API_BASE_URL"); envURL != "" {
		config.BaseURL = envURL
	}
	if envKey := os.Getenv("API_KEY"); envKey != "" {
		config.ApiKey = envKey
	}
	if envAccount := os.Getenv("ACCOUNT_ID"); envAccount != "" {
		config.AccountId = envAccount
	}
	if envUser := os.Getenv("USER_ID"); envUser != "" {
		config.UserId = envUser
	}

	fmt.Println("=== JobModel Client Library Test ===")
	fmt.Printf("API URL: %s\n", config.BaseURL)
	fmt.Printf("Account ID: %s\n", config.AccountId)
	fmt.Println()

	client := NewClientWithDefaults(config.BaseURL, config.ApiKey)

	// Test 1: Create a job
	fmt.Println("1. Testing job creation...")
	job, err := testCreateJob(client, config.AccountId, config.UserId)
	if err != nil {
		log.Printf("❌ Create job failed: %v", err)
		return
	}
	fmt.Printf("✅ Job created successfully: %s\n", job.Title)
	fmt.Printf("   Partition Key: %s\n", job.PartitionKey)
	fmt.Printf("   Range Key: %s\n", job.RangeKey)
	fmt.Println()

	// Test 2: Fetch the created job
	fmt.Println("2. Testing single job fetch...")
	fetchedJob, err := testFetchSingleJob(client, job.PartitionKey, job.RangeKey)
	if err != nil {
		log.Printf("❌ Fetch job failed: %v", err)
	} else {
		fmt.Printf("✅ Job fetched successfully: %s\n", fetchedJob.Title)
	}
	fmt.Println()

	// Test 3: Fetch all jobs
	fmt.Println("3. Testing fetch all jobs...")
	allJobs, err := testFetchAllJobs(client, config.AccountId)
	if err != nil {
		log.Printf("❌ Fetch all jobs failed: %v", err)
	} else {
		fmt.Printf("✅ Fetched %d jobs successfully\n", len(allJobs.Results))
	}
	fmt.Println()

	// Test 4: Update the job
	fmt.Println("5. Testing job update...")
	updatedJob, err := testUpdateJob(client, job, config.UserId)
	if err != nil {
		log.Printf("❌ Update job failed: %v", err)
	} else {
		fmt.Printf("✅ Job updated successfully: %v\n", updatedJob)
	}
	fmt.Println()

	// Test 5: Delete the job (cleanup)
	fmt.Println("7. Testing job deletion (cleanup)...")
	err = testDeleteJob(client, job.PartitionKey, job.RangeKey)
	if err != nil {
		log.Printf("❌ Delete job failed: %v", err)
	} else {
		fmt.Printf("✅ Job deleted successfully\n")
	}
	fmt.Println("\n=== Test Complete ===")
}

func testCreateJob(client *Client, accountId, userId string) (*model.JobModel, error) {
	req := request.JobModelCreateRequest{
		AccountId:           accountId,
		UserId:              userId,
		Title:               fmt.Sprintf("Test Job - %d", time.Now().Unix()),
		Description:         "This is a test job created by the Go client library",
		Department:          "Engineering",
		Category:            "Software Development",
		Location:            "Remote",
		JobType:             "full-time",
		SalaryMin:           80000,
		SalaryMax:           120000,
		Requirements:        "Go programming, AWS experience",
		Responsibilities:    "Develop and maintain backend services",
		Benefits:            "Health, dental, 401k",
		PostedBy:            "test@example.com",
		ApplicationDeadline: "2024-12-31",
		ExperienceLevel:     "mid",
		RemoteAllowed:       true,
	}

	return client.CreateJob(req)
}

func testFetchSingleJob(client *Client, partitionKey, rangeKey string) (*model.JobModel, error) {
	return client.FetchJobSimple(partitionKey, rangeKey)
}

func testFetchAllJobs(client *Client, accountId string) (*model.JobModelsResponse, error) {
	req := request.JobModelFetchAllRequest{
		AccountId: accountId,
		Limit:     10,
	}
	return client.FetchAllJobs(req)
}

func testUpdateJob(client *Client, job *model.JobModel, userId string) (*bool, error) {
	job.Title = job.Title + " (Updated)"
	job.SalaryMax = job.SalaryMax + 10000

	req := request.JobModelUpdateRequest{
		Job:    *job,
		UserId: userId,
	}

	return client.UpdateJob(req)
}

func testDeleteJob(client *Client, partitionKey, rangeKey string) error {
	return client.DeleteJobSimple(partitionKey, rangeKey)
}

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Println(string(b))
}
