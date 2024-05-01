package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/whiteSpaceController", whiteSpaceController)
	fmt.Println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func whiteSpaceController(w http.ResponseWriter, r *http.Request) {
	// Request body's text
	body := r.Body
	defer body.Close()

	// Read the text from the body
	text, err := io.ReadAll(body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Check for whitespace
	if strings.ContainsAny(string(text), " \t\n\r") {
		fmt.Fprintf(w, "The text contains whitespace")
	} else {
		fmt.Fprintf(w, "The text does not contain whitespace")
	}
}
