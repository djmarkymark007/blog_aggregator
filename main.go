package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	serverHandler := http.NewServeMux()
	serverHandler.HandleFunc("GET /v1/readiness", readiness)
	serverHandler.HandleFunc("GET /v1/err", errTest)

	server := http.Server{Handler: serverHandler, Addr: ":" + port}
	fmt.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
