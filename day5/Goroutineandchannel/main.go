package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Task 1: countNumbers counts from 1 to 5 and sends each number as a string to the 'out' channel.
// Send the current number as a formatted string to the channel
// Sleep to simulate work and allow interleaving with other goroutines
func countNumbers(out chan<- string) {
	for i := 1; i <= 5; i++ {
		out <- fmt.Sprintf("Number: %d", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Task 2: printLetters sends letters A through E as strings to the 'out' channel.
// Send each letter to the channel
// Sleep longer than countNumbers to illustrate asynchronous behavior

func printLetters(out chan<- string) {
	letters := []rune{'A', 'B', 'C', 'D', 'E'}
	for _, letter := range letters {

		out <- fmt.Sprintf("Letter: %c", letter)
		time.Sleep(700 * time.Millisecond)
	}
}

// Task 3: fetchURL performs an HTTP GET request and sends the first 200 bytes of the response body to 'out'.
// 45-Ensure the response body is closed when function exits
// 47-Read entire response body
// Send truncated response body (first 200 bytes) to avoid flooding output
func fetchURL(out chan<- string) {
	resp, err := http.Get("https://www.youtube.com")
	if err != nil {

		out <- fmt.Sprintf("Error fetching URL: %v", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		out <- fmt.Sprintf("Error reading response: %v", err)
		return
	}
	out <- "Fetched data from given URL : " + string(body[:200]) + "..."
}

// 61-Channel to send strings from all goroutines to the main goroutine
// 63-Channel to signal completion of each goroutine
// Launch countNumbers, prinetletters and fetchurl in different goroutine
// Signal completion by sending true into 'done'
// 82-Collector goroutine to wait for all tasks to finish before closing 'out'.Wait for 3 completion signals from done channel.
func main() {

	out := make(chan string)

	done := make(chan bool)

	go func() {
		countNumbers(out)

		done <- true
	}()

	go func() {
		printLetters(out)
		done <- true
	}()

	go func() {
		fetchURL(out)
		done <- true
	}()

	go func() {

		for i := 0; i < 3; i++ {
			<-done
		}
		close(out)
	}()

	for msg := range out {
		fmt.Println(msg)
	}

	fmt.Println("✅ All tasks completed.")
}
