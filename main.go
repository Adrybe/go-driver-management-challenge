package main

import (
	"log"
	"net/http"
)

func main() {
	// Init Router
	port, app := processor.Setup()
	// Database Connection
	_, _ = repository.Connect()
	// Start Server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(port, app))
}
