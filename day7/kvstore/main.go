package main

import (
	"fmt"
	"log"
	"net/http"

	"kvstore/api"
	"kvstore/pkg/logger"
)

func main() {
	logger.Init()
	log.Println("Starting server on :8080")

	http.HandleFunc("/set", api.SetHandler)
	http.HandleFunc("/get", api.GetHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
