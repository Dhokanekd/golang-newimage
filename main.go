package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// The handler function for the web server
func handler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request to the console
	log.Printf("Received request from %s on path: %s\n", r.RemoteAddr, r.URL.Path)

	// Write the response back to the client
	fmt.Fprintf(w, "Hello from the Go Application! You requested: %s\n", r.URL.Path)
}

func main() {
	// The port to listen on. OpenShift often sets this via environment variables.
	// Default to "8080" if the PORT environment variable is not set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	// Set up the handler for all requests ("/")
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server on port %s\n", port)
	
	// Start the server and log any error that prevents it from starting
	// http.ListenAndServe is a blocking function.
	log.Fatal(http.ListenAndServe(addr, nil))
}
