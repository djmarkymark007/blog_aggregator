package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/djmarkymark007/blog_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func getApiKey(r *http.Request) string {
	Key := r.Header.Get("Authorization")
	Key = strings.TrimLeft(Key, "ApiKey ")
	return Key
}

// Should this be in a different file?
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
