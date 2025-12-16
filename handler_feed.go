package main

import (
	"context"
	"fmt"
	"time"

	"github.com/glebson1988/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	userName := s.cfg.CurrentUserName
	if userName == "" {
		return fmt.Errorf("no current user, please login or register first")
	}

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	now := time.Now().UTC()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:   %v\n", feed.ID)
	fmt.Printf(" * Name: %v\n", feed.Name)
	fmt.Printf(" * URL:  %v\n", feed.Url)

	return nil
}
