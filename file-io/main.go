package main

import (
	"fmt"
	"os"
)

func main() {
	// Define a sample file path
	filePath := "sample.txt"

	// Create a sample file
	content := []byte("Hello, File I/O in Go!")

	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	fmt.Println("File created successfully.")

	// Read the contents of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	fmt.Println("File content:", string(data))

	// Open a file for appending
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Append content to the file
	newContent := []byte("\nAppended line.")
	_, err = file.Write(newContent)
	if err != nil {
		fmt.Println("Error appending to the file:", err)
		return
	}
	fmt.Println("Content appended successfully.")

	// Read the updated file contents
	data, err = os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the updated file:", err)
		return
	}
	fmt.Println("Updated file content:", string(data))
}
