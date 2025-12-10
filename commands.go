package main

import "fmt"

type Commands struct {
	registeredCommands map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registeredCommands[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command %q not found", cmd.Name)
	}
	return handler(s, cmd)
}
