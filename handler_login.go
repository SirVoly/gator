package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	//Check if user exists
	_, err := s.db.GetUser(
		context.Background(),
		username,
	)
	if err != nil {
		return fmt.Errorf("user %s does not exist", username)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set: %s\n", username)
	return nil
}
