package main

import (
	"fmt"
	"net/http"
	"time"
)

func startWorkers(workerId int, createJobsForWorkers chan int, resultsFromWorkers chan string) {
	for range createJobsForWorkers {
		resp, err := http.Get("http://localhost:8080/api/v1/get-data")
		if err != nil {
			resultsFromWorkers <- fmt.Sprintf("Worker %d: Error: %s", workerId, err)
		} else {
			resultsFromWorkers <- fmt.Sprintf("Worker %d: Response: %s", workerId, resp.Status)
			resp.Body.Close()
		}
	}
}

func main() {
	numberOfConcurrentRequests := 1000
	numberOfWorkers := 100

	createJobsForWorkers := make(chan int, numberOfConcurrentRequests)
	resultsFromWorkers := make(chan string, numberOfConcurrentRequests)

	// We will now start the workers
	for worker := 1; worker <= numberOfWorkers; worker++ {
		go startWorkers(worker, createJobsForWorkers, resultsFromWorkers)
	}

	// Send the jobs to the workers
	start := time.Now()
	for job := 1; job <= numberOfConcurrentRequests; job++ {
		createJobsForWorkers <- job
	}
	close(createJobsForWorkers)

	// Get the results from the workers
	for result := 1; result <= numberOfConcurrentRequests; result++ {
		fmt.Println(<-resultsFromWorkers)
	}

	fmt.Printf("Time taken: %s\n", time.Since(start))
}
