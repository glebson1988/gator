package main

import "fmt"

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
