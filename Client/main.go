package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendConcurrentRequest(client http.Client, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Request to %s completed with status code: %d\n", url, resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	loadBalancerUrl := "http://localhost:8080/api/v1/load-balancer"

	numberOfConcurrentRequests := 1000

	wg.Add(numberOfConcurrentRequests)

	startTime := time.Now()

	for i := 0; i < numberOfConcurrentRequests; i++ {
		go sendConcurrentRequest(client, loadBalancerUrl, &wg)
	}

	wg.Wait()

	elaspedTime := time.Since(startTime)

	fmt.Printf("Completed %d requests in %v time: \n", numberOfConcurrentRequests, elaspedTime)
}
