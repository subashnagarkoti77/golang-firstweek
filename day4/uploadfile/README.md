#  Go File Upload API

This simple Go web server allows users to upload a text file via an HTTP POST request. It saves the uploaded file to disk.

---

##  Features

- Accepts file uploads via `/upload` endpoint.
- Saves uploaded files locally as `uploaded_<filename>`.
- Responds with a success message. 


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

