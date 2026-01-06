package main

import (
	"errors";
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	CommandToHandler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	commandHandler, ok := c.CommandToHandler[cmd.Name]
	if !ok {
		return errors.New(fmt.Sprintf("command %s not found", cmd.Name))
	}

	return commandHandler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.CommandToHandler[name] = f
}

func generateCommands() commands {
	cmds := commands{
		CommandToHandler: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	return cmds
}
