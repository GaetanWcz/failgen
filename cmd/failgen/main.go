package main

import (
	"log"
	"net/http"
	"os/exec"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "failgen/docs"
	"failgen/internal/handler"
)

func init() {
	cmd := exec.Command("swag", "init", "-g", "cmd/failgen/main.go", "-o", "docs")
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to generate Swagger docs: %v", err)
	}
}

func main() {
	http.HandleFunc("/fail", handler.FailHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("[FailGen] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
