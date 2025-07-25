package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// 31 - This runs at the end of the function automatically.Use dst to write data here.
// When function ends, dst.Close() is called automatically
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Only Post Method is Allowed", http.StatusMethodNotAllowed)
		return

	}
	var dst *os.File
	var file multipart.File
	var header *multipart.FileHeader
	var err error
	file, header, err = r.FormFile("file")
	if err != nil {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err = os.Create("uploaded_" + header.Filename)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", header.Filename)

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
