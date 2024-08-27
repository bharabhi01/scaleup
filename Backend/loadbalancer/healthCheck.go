package loadbalancer

import (
	"fmt"
	"net/http"
	"time"
	"log"
)


// Periodically perform health checks of each server
func(lb *LoadBalancer) PerformHealthChecks(interval *time.Duration) {
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			lb.checkServerHealth()
		}
	}
}

// Iterate over the servers and update their heath status
func(lb *LoadBalancer) checkServerHealth() {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	for _,server := range lb.Servers {
		go func(s *Server) {
			url := "http://" + s.Address + "/health"
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK{
				s.Mutex.Lock()
				s.Alive = false
				s.Mutex.Unlock()
				log.Printf("Server %s is down", s.Address)
			} else {
				s.Mutex.Lock()
				s.Alive = true
				s.Mutex.Unlock()
				log.Printf("Server %s is up", s.Address)
			}	
			if resp != nil {
				resp.Body.Close()
			}
		} (server)	
	}
}