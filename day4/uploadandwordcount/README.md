#  Go File Upload and Word Count API

This simple Go web server allows users to upload a text file via an HTTP POST request. It saves the uploaded file to disk and then calculates the total number of words in the file.

---

##  Features

- Accepts file uploads via `/upload` endpoint.
- Saves uploaded files locally as `uploaded_<filename>`.
- Reads and processes the saved file.
- Returns the total number of words in the uploaded file.

---

###  How to Run

Make sure you have Go installed. Then:


```bash

go run main.go

```

### After the program succesfully run it will show message as:
Server is running at http://localhost:8080"

----Now, In new terminal----

curl -F "file=@test.txt" http://localhost:8080/upload


