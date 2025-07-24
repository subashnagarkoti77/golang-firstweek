package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// This function handles requests to /hello. Only allow GET requests (e.g. from browser or curl)
func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello, World!")

}

// This function handles requests to /time.
// Create a map (like a dictionary) to hold the current time.
// Set the content type of the response to JSON.
// Convert the map to JSON and write it to the response
func timeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	currentTime := map[string]string{
		"time": time.Now().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentTime)
}

// Tell the server to run helloHandler when /hello is requested, and timeHandler when /time is requested.
// start the server port on 8080
func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
