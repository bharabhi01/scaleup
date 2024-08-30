package backendserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// StartServer starts the backend server on the specified port
func StartServer(port string) {
	// Define a simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Define a simple handler for general requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server on port %s at %s\n", port, time.Now().Format(time.RFC3339))
	})

	// Start the server
	log.Printf("Starting backend server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
