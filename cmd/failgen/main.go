package main

import (
	"log"
	"net/http"

	"failgen/internal/handler"
)

func main() {
	http.HandleFunc("/fail", handler.FailHandler)
	log.Println("[FailGen] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}