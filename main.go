package main

import (
	"fmt"
	"net/http"
	"os"
	
	ourcode "main.go/handlers"
)

func main() {
	http.HandleFunc("/css/", ourcode.CssHandler)
	http.HandleFunc("/", ourcode.HomeHandler)
	http.HandleFunc("/ascii-art", ourcode.AsciiArtHandler)
	fmt.Println("Server starting on :8080\nVisit: http://127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
