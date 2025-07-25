package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// 27 - Get uploaded file from form
// 32 - This runs at the end of the function automatically.Use dst to write data here.
// When function ends, dst.Close() is called automatically
// 34 - save uplaoded file to disk
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

	filename := "uploaded_" + header.Filename
	dst, err = os.Create(filename)
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
	wordCount, err := countWordsInFile(filename)
	if err != nil {
		http.Error(w, "Failed to process file for word count", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File '%s' uploaded successfully.\nTotal number of words: %d\n", header.Filename, wordCount)

}

func countWordsInFile(filepath string) (int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0, err
	}

	text := string(data)
	words := strings.Fields(text)
	return len(words), nil
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
