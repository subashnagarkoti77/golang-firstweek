#  In-Memory Key-Value Store API in GOLANG

This is a simple HTTP-based key-value store built using Go. It stores data in memory and provides endpoints to set and get key-value pairs.

---

##  Features

-  Set a key-value pair using `POST /set`
-  Get the value of a key using `GET /get?key=your_key`
-  Data is stored in memory using a Go `map[string]string`

---

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
  -d '{"key":"username", "value":"subash"}'

-- curl "http://localhost:8080/get?key=username" 

