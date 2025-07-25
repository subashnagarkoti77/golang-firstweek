package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Simple in-memory map to store key-value pairs
// Structure to read JSON from POST /set
var store = make(map[string]string)

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// handle POST /set rquest
//28-Decode the json body into SetRequest structure
//35-store the key pair value in map

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method", http.StatusMethodNotAllowed)
		return
	}

	var req SetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Key == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	store[req.Key] = req.Value
	fmt.Fprintf(w, "Saved key: %s\n", req.Key)
}

// 47-Get the 'key' from the URL query
// 53-Look up the key in the map
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Use GET method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
		return
	}

	value, ok := store[key]
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Value: %s\n", value)
}

// Register the route /set with setHandler function and getHandler Function
func main() {
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
