package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadConfig() *Config {
	return &Config{
		Port:                   getIntEnv("Port", 8080),
		Servers:                strings.Split(getEnv("SERVERS", "localhost:9001, localhost:9002"), ","),
		LoadBalancingAlgorithm: getEnv("LB_ALGORITHM", "roundrobin"),
		HealthCheckInterval:    time.Second * getDurationEnv("HEALTH_CHECK_INTERVAL", 10),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if durationValue, err := time.ParseDuration(value + "s"); err == nil {
			return durationValue
		}
	}

	return defaultValue
}
