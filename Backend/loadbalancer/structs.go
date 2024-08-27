package loadbalancer

import "sync"

// Server represents a server with its address and health status
type Server struct {
	Address string
	Alive   bool
	Mutex   sync.RWMutex
}

// LoadBalancer represents the load balancer with a list of servers and the current server index
type LoadBalancer struct {
	Servers []*Server
	Current int
	Mutex   sync.Mutex
}
