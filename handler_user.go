package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/djmarkymark007/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func createUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("failed to decode body. %s", err)
		respondWithError(w, 400, "failed to convert body to json")
		return
	}

	uuidString := uuid.New().String()
	if uuidString == "" {
		log.Printf("failed to create uuid. %s", err)
		respondWithError(w, 500, "failed to create user")
		return
	}

	query := database.CreateUserParams{
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		ID:        uuid.New(),
	}

	user, err := config.DB.CreateUser(r.Context(), query)
	if err != nil {
		log.Printf("couldn't create user. %s", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create user")
		return
	}

	respondWithJson(w, 200, databaseUsertoUser(user))
}

func getApiKey(r *http.Request) string {
	apiKey := r.Header.Values("Authorization")[0]
	apiKey = strings.TrimLeft(apiKey, "ApiKey ")
	return apiKey
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	apiKey := getApiKey(r)
	if len(apiKey) != 64 {
		respondWithError(w, 400, "Invalid api key")
		return
	}
	user, err := config.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		log.Printf("couldn't get user by api key. %s", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't get user")
		return
	}

	respondWithJson(w, 200, databaseUsertoUser(user))
}
