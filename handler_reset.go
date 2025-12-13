package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}

	return nil
}
