package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/bharabhi01/scaleup/Backend/config"
	"github.com/bharabhi01/scaleup/Backend/loadbalancer"
)

func main() {
	cfg := config.LoadConfig()

	lb := loadbalancer.NewLoadBalancer(cfg.Servers)

	// Set up health checks
	go func() {
		for {
			lb.PerformHealthChecks(cfg.HealthCheckInterval)
			time.Sleep(cfg.HealthCheckInterval)
		}
	}()

	// Handle incoming requests and forward them to the backend servers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server := lb.GetServer(cfg.LoadBalancingAlgorithm)
		if server == nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		proxyUrl := fmt.Sprintf("http://%s %s", server.Address, r.URL.Path)
		resp, err := http.Get(proxyUrl)
		if err != nil {
			http.Error(w, "Error proxying request", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		// Now we will copy the response from the backend server to the client
		for k, v := range resp.Header {
			w.Header()[k] = v
		}

		w.WriteHeader(resp.StatusCode)
		_, err = http.Copy(w, resp.Body)
		if err != nil {
			log.Printf("Error copying response body: %v", err)
		}
	})

	// Start the server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Load Balancing started on %s using %s algorithm", addr, cfg.LoadBalancingAlgorithm)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
