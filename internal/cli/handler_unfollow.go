package cli

import (
	"context"
	"fmt"

	"github.com/SirVoly/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feedURL>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.RemoveFeedFollow(
		context.Background(),
		url,
	)

	return err
}