package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("login takes exactly one argument: username")
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("User has been set: %s\n", cmd.Args[0])
	return nil
}
