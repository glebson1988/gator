package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/glebson1988/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	limit := 0
	if len(cmd.Args) == 0 {
		limit = 2
	} else {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid conversion string to int: %w", err)
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time)
		}
		fmt.Println("---") // separator between posts
	}

	return nil
}
