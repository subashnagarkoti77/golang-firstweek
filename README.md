This repository contains a step-by-step learning journey to explore Go programming fundamentals by building practical projects each day. The project gradually grows from simple programs to a working in-memory key-value REST API.

Over 7 days, I covered:

- Basic syntax and core programming concepts
- Building command-line tools
- Creating simple REST APIs
- Handling files and JSON data
- Using goroutines and channels for concurrency
- Implementing a basic in-memory key-value store with HTTP interface
- Improving code structure, adding logging, and writing unit tests

## Day-wise Breakdown

# Day 1: Go Basics
- Installed Go and configured VS Code environment.
- Wrote a "Hello World" program.
- Learned variables, loops, functions, error handling, and conditionals.
# Day 2: CLI Basics
- Built a calculator CLI supporting addition, subtraction, multiplication, and division.
- Practiced reading user input via fmt.Scanln.
- Explored Go package structure and modules.
# Day 3: Simple REST API
- Created an HTTP server with endpoints:
        /hello returns "Hello World"
        /time returns the current server time.
- Learned basics of net/http package.
# Day 4: File Handling + JSON
- Wrote a program to Read text files and count words.
- Wrote a program to Upload file and word count of that file.
- Practiced JSON decoding and file processing.
# Day 5: Goroutines and Channels
- Wrote a program to run multiple tasks concurrently using goroutines and channels.
- Examples included counting numbers, printing letters and Fetch data from URL concurrently.
- Understood Go’s concurrency model.
# Day 6: In-memory Key-Value Store API
- Implemented /set (POST) and /get (GET) endpoints.
- Used Go map to store data in-memory.
- Built a simple REST API demonstrating basic CRUD concepts.
# Day 7: Code Cleanup & Presentation
- Refactored code into a clean folder structure (api/, internal/store/, pkg/logger/).
- Added logging for better visibility using a centralized logger package.
- Wrote unit tests for store functions using Go’s testing package.
- Prepared this README summarizing the project.