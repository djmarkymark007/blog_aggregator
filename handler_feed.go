package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/djmarkymark007/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameter{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	feedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.Url,
	}

	feed, err := config.DB.CreateFeed(r.Context(), feedArgs)
	if err != nil {
		log.Printf("couldn't create feed. %s", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed")
		return
	}
	respondWithJson(w, 200, databaseFeedtoFeed(feed))
}

// TODO(Mark): set a return limit?
func getFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := config.DB.GetFeeds(r.Context())
	if err != nil {
		log.Printf("couldn't get feeds. %s", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't get feeds")
		return
	}

	respondWithJson(w, 200, databaseFeedstoFeeds(feeds))
}
