package main

import (
	"context"
	"fmt"

	"github.com/glebson1988/gator/internal/database"
)

type Commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *Commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *Commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command %q not found", cmd.Name)
	}
	return handler(s, cmd)
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
