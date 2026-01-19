package cli

import (
	"context"
	"fmt"

	"github.com/SirVoly/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	
	return func(s *state, cmd command) error {
		username := s.cfg.GetUser()
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("user %s was not found", username)
		}

		return handler(s, cmd, user)
	}
}