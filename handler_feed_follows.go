package main

import (
	"context"
	"fmt"

	"github.com/glebson1988/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	userName := s.cfg.CurrentUserName
	if userName == "" {
		return fmt.Errorf("no current user, please login or register first")
	}

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created successfully:")
	fmt.Printf(" * Feed name:   %v\n", feedFollow.FeedName)
	fmt.Printf(" * User name: %v\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	userName := s.cfg.CurrentUserName
	if userName == "" {
		return fmt.Errorf("no current user, please login or register first")
	}

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range feedFollows {
		fmt.Println(f.FeedName)
	}

	return nil
}
