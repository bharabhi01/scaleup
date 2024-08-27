package loadbalancer

import (
	"math/rand"
	"time"
)

func NewLoadBalancer(serverAddresses []string) *LoadBalancer {
	servers := make([]*Server, len(serverAddresses))
	for i, address := range serverAddresses {
		servers[i] = &Server{
			Address: address,
			Alive:   true,
		}
	}

	return &LoadBalancer{
		Servers: servers,
	}
}

// Round Robin Algorithm
func (lb *LoadBalancer) RoundRobin() *Server {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	server := lb.Servers[lb.Current]
	lb.Current = (lb.Current + 1) % len(lb.Servers)
	return server
}

// Random Selection Algorithm
func (lb *LoadBalancer) RandomSelection() *Server {
	rand.Seed(time.Now().UnixNano())
	return lb.Servers[rand.Intn(len(lb.Servers))]
}

// Get Server based on the algorithm
func(lb *LoadBalancer) GetServer(algorithm string) *Server {
	switch algorithm {
		case "roundrobin":
			return lb.RoundRobin()
	
		case "random" :
			return lb.RandomSelection()

		default:
			return lb.RoundRobin()
	}
}
