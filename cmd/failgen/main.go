package main

import (
	"log"
	"net/http"

	"failgen/internal/handler"
)

func main() {
	http.HandleFunc("/fail", handler.FailHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	log.Println("[FailGen] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}