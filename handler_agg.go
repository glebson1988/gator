package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/glebson1988/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("error scraping feeds: %v\n", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return fmt.Errorf("could not get next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't set feed as last fetched %w", err)
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		publishedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("error parsing date for RSSItem %s: %v\n", item.Title, err)
			continue
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedTime, Valid: true},
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue // ignore duplicates
			}
			fmt.Printf("error creating post: %v\n", err)
		}

	}

	return nil
}
