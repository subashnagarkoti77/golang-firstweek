package api

import (
	"encoding/json"
	"kvstore/internal/store"
	"net/http"
)

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
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

	store.Set(req.Key, req.Value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Saved key: " + req.Key))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Use GET method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' parameter", http.StatusBadRequest)
		return
	}

	value, ok := store.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Value: " + value))
}
