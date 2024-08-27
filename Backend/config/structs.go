package config

import "time"

type Config struct {
	Port                   int
	Servers                []string
	LoadBalancingAlgorithm string
	HealthCheckInterval    time.Duration
}
