package main

import (
	"log"
	"matrix-operations/internal/http_operations"
	"net/http"
)

func main() {
	handleHttpOperations()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server on port 8080")
	}

	log.Print("Listening on port 8080")
}

func handleHttpOperations() {
	ops := http_operations.GetHttpOperations()

	http.HandleFunc("/echo", ops.Echo)
	http.HandleFunc("/invert", ops.Invert)
	http.HandleFunc("/flatten", ops.Flatten)
	http.HandleFunc("/sum", ops.Sum)
	http.HandleFunc("/multiply", ops.Multiply)
}
