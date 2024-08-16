package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCount uint64

func main() {
	http.HandleFunc("/", handleRequest)

	go printStats()

	fmt.Println("Load Balancer is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Increment the request count
	atomic.AddUint64(&requestCount, 1)

	// Log the incoming request
	log.Printf("Request processed: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

// Periodically print the total number of requests processed
func printStats() {
	for {
		time.Sleep(5 * time.Second)
		count := atomic.LoadUint64(&requestCount)
		log.Printf("Total requests processed: %d", count)
	}
}
