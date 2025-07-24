# Simple Go HTTP REST API

This is a basic REST API written in Go using the standard `net/http` package.  
It includes two simple endpoints: `/hello` and `/time`.

After the program succesfully run it will show message as:
Server is running at http://localhost:8080"

----Now, In new terminal----

`curl localhost:8080/hello`: Returns a plain text "Hello, World!" message.
`curl localhost:8080/time`: Returns the current server time in JSON format.

```bash
go run /day3/hello+time-api/main.go

