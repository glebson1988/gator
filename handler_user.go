package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("No arguments")
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("You were successfully logged in")
	return nil
}
