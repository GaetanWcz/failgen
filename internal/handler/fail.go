package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"failgen/internal/response"
)

func FailHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	statusCode := 500
	if status := query.Get("status"); status != "" {
		if code, err := strconv.Atoi(status); err == nil {
			statusCode = code
		}
	}

	delayMs := 0
	if delay := query.Get("delay"); delay != "" {
		if ms, err := strconv.Atoi(delay); err == nil {
			delayMs = ms
		}
	}

	message := query.Get("message")
	if message == "" {
		message = http.StatusText(statusCode)
	}

	if delayMs > 0 {
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	json.NewEncoder(w).Encode(response.ErrorResponse{
		Status:  statusCode,
		Message: message,
		Details: "Simulated error by FailGen",
	})
}
