package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ryutah/learn-k8s/internal/frontend"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), frontend.Handler()))
}
