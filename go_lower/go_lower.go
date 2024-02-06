package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Check if a file name is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run yourprogram.go /path/to/filename.txt")
		os.Exit(1)
	}

	// Get the filename from the command-line argument
	filename := os.Args[1]

	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Convert uppercase to lowercase
	modifiedContent := strings.ToLower(string(content))

	// Write the modified content back to the file
	err = ioutil.WriteFile(filename, []byte(modifiedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("Uppercase to lowercase conversion complete.")
}

//to run the programm 'go run go_lower 'file_path' on terminal
// which will make all the uppercase letter to lowercase
