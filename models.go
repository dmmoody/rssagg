package main

import (
	"github.com/dmmoody/rssagg/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(databaseUser database.User) User {
	return User{
		ID:        databaseUser.ID,
		Name:      databaseUser.Name,
		ApiKey:    databaseUser.ApiKey,
		CreatedAt: databaseUser.CreatedAt,
		UpdatedAt: databaseUser.UpdatedAt,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(databaseFeed database.Feed) Feed {
	return Feed{
		ID:        databaseFeed.ID,
		Name:      databaseFeed.Name,
		Url:       databaseFeed.Url,
		UserID:    databaseFeed.UserID,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
	}
}

func databaseFeedToFeeds(databaseFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range databaseFeed {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(databaseFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        databaseFeed.ID,
		UserID:    databaseFeed.UserID,
		FeedID:    databaseFeed.ID,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(databaseFeed []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, feedFollow := range databaseFeed {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePostToPost(databasePost database.Post) Post {
	var description *string
	if databasePost.Description.Valid {
		description = &databasePost.Description.String
	}

	return Post{
		ID:          databasePost.ID,
		FeedID:      databasePost.FeedID,
		Title:       databasePost.Title,
		Url:         databasePost.Url,
		Description: description,
		PublishedAt: databasePost.PublishedAt,
		CreatedAt:   databasePost.CreatedAt,
		UpdatedAt:   databasePost.UpdatedAt,
	}
}

func databasePostsToPosts(databasePosts []database.Post) []Post {
	posts := []Post{}
	for _, post := range databasePosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
