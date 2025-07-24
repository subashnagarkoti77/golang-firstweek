package main

import (
	"fmt"
	"os"
	"strings"
)

var data []byte
var text string
var words []string
var wordcount int
var err error

// read the contents of file
// 26-convert to string and split into words
// 27-spilts on spaces, tabs, and newlines
func main() {

	data, err = os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	text = string(data)
	words = strings.Fields(text)

	wordcount = len(words)

	fmt.Println("Total number of words:", wordcount)
}
