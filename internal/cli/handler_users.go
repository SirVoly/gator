package cli

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(
		context.Background(),
	)
	if err != nil {
		return fmt.Errorf("couldn't get all users: %w", err)
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf(" (current)")
		}
		fmt.Println()
	}
	return nil
}
