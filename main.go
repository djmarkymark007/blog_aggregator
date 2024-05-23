package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/djmarkymark007/blog_aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

var config apiConfig

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbURL := os.Getenv("PROTOCOL")
	if dbURL == "" {
		log.Fatal("protocol environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	config = apiConfig{
		DB: database.New(db),
	}

	serverHandler := http.NewServeMux()
	serverHandler.HandleFunc("GET /v1/readiness", readiness)
	serverHandler.HandleFunc("GET /v1/err", errTest)
	serverHandler.HandleFunc("POST /v1/users", createUsers)

	server := http.Server{Handler: serverHandler, Addr: ":" + port}

	fmt.Printf("Starting server on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
