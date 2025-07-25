#  Go File Upload and Word Count API

This simple Go web server allows users to upload a text file via an HTTP POST request. It saves the uploaded file to disk and then calculates the total number of words in the file.

---

##  Features

- Accepts file uploads via `/upload` endpoint.
- Saves uploaded files locally as `uploaded_<filename>`.
- Reads and processes the saved file.
- Returns the total number of words in the uploaded file.

---



```bash
cd day4/uploadandwordcount
go run main.go
