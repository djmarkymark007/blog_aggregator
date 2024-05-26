package main

import (
	"time"

	"github.com/djmarkymark007/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUsertoUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreateAt:  user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreateAt:  feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserId:    feed.UserID,
	}
}

func databaseFeedstoFeeds(feeds []database.Feed) []Feed {
	res := make([]Feed, len(feeds))
	for index, feed := range feeds {
		res[index] = databaseFeedtoFeed(feed)
	}
	return res
}
