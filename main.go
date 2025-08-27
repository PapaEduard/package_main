package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Metrics struct {
	Status string `json:"status"`
	Data   struct {
		Metrics struct {
			CPU     float64 `json:"cpu_usage"`
			Memory  float64 `json:"memory_usage"`
			Threads int     `json:"thread_count"`
		} `json:"metrics"`
	} `json:"data"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		metrics := Metrics{Status: "success"}
		metrics.Data.Metrics.CPU = rand.Float64()
		metrics.Data.Metrics.Memory = rand.Float64() * 100
		metrics.Data.Metrics.Threads = rand.Intn(500) + 1

		json.NewEncoder(w).Encode(metrics)
	})

	http.ListenAndServe(":9200", nil)
}
