package main

import (
	"encoding/json"
	"net/http"
)

type jasonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jasonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	// struct -> jason string
	out, _ := json.MarshalIndent(payload, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
