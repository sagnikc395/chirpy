package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	mux := http.NewServeMux()
	server := &http.Server{Handler: mux, Addr: PORT}

	fmt.Printf("server started at http://localhost:%s\n", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}
