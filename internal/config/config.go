package config

import (
	"os"
	"strings"
)

type Config struct {
	WorkerNodes  []string
	WorkerAPIKey string
}

func Load() Config {
	return Config{
		WorkerNodes:  parseList(os.Getenv("WORKER_NODES")),
		WorkerAPIKey: os.Getenv("WORKER_API_KEY"),
	}
}

func parseList(value string) []string {
	var items []string
	for item := range strings.SplitSeq(value, ",") {
		if item = strings.TrimSpace(item); item != "" {
			items = append(items, item)
		}
	}
	return items
}
