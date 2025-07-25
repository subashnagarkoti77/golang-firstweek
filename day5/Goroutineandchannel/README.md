### Concurrent Go Program: Numbers, Letters, and HTTP Fetch

This Go program demonstrates how to use goroutines and channels to run multiple tasks concurrently:

### What It Does

The program performs three tasks at the same time:

1. Counts numbers from 1 to 5
2. Prints letters from A to E
3. Fetches data from a URL (`https://www.youtube.com`) and prints the first 200 bytes of the response

Each task runs in its own goroutine and sends messages to a shared output channel, which the main function listens to and prints.

---

###  Concurrency

- Uses Go's goroutine to perform tasks concurrently
- Uses channels to:
  - Send results back to the main function
  - Signal task completion
- A separate goroutine waits for all tasks to finish before closing the output channel

---

###  How to Run

Make sure you have Go installed. Then:

```bash

go run main.go
