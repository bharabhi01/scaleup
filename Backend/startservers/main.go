package main

import (
	"log"
	"os"

	"github.com/bharabhi01/scaleup/backend/backendserver"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Port number is required as an argument")
	}
	port := os.Args[1]

	// Start the backend server on the specified port
	backendserver.StartServer(port)
}
