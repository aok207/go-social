package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Healthy bool `json:"healthy"`
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HealthResponse{
		Healthy: true,
	})
}
