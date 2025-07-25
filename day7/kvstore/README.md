# KVStore API

A simple in-memory Key-Value Store API written in Go. This project demonstrates basic REST endpoints, modular project structure, logging, and unit testing using Go's standard library.

---

##  Features

- `POST /set` — Store a key-value pair
- `GET /get?key=yourKey` — Retrieve value by key
- Simple in-memory storage (map)
- Modular folder structure
- Logging using the `log` package
- Unit tests for store logic

---
## How to do Unit testing
This test calls the Set and Get functions directly to verify they work correctly without involving the whole application (e.g., no HTTP requests, no database).

In bash:

go test ./internal/store

-----

##  How to Run

- Make sure you have installed Go version 1.18 or later

   ```bash
  go run main.go
  
  ```

  ### After the program succesfully run it will show message as:
Server is running at http://localhost:8080"

----Now, In new terminal----

-- curl -X POST http://localhost:8080/set \
  -H "Content-Type: application/json" \
  -d '{"key":"yourkey", "value":"yourvalue"}'

-- curl "http://localhost:8080/get?key=yourkey" 

