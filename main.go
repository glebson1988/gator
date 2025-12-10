package main

import (
	"log"
	"os"

	"github.com/glebson1988/gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &State{
		Cfg: &cfg,
	}

	cmds := Commands{
		registeredCommands: make(map[string]func(*State, Command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Too few arguments")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
