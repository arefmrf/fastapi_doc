package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<h1>Hello, World!</h1>")
	//_, err := w.Write([]byte("<h1>Hello, World!</h1>"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	serverAddr := ":3000"
	fmt.Printf("Starting server on %s...\n", serverAddr)

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
