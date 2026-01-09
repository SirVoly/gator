package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	id := uuid.New()
	t := time.Now()

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        id,
			CreatedAt: t,
			UpdatedAt: t,
			Name:      username,
		},
	)
	if err != nil {
		return fmt.Errorf("user %s already exists", username)
	}

	s.cfg.SetUser(username)

	fmt.Printf("User was created: %s\n", username)

	// User Log
	fmt.Printf("User: %s\n", user.Name)
	fmt.Printf("\tid: %s\n", user.ID)
	fmt.Printf("\tcreated_at: %s\n", user.CreatedAt)
	fmt.Printf("\tupdated_at: %s\n", user.UpdatedAt)

	return nil
}
