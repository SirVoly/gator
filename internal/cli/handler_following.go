package cli

import (
	"context"
	"fmt"

	"github.com/SirVoly/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	followedfeeds, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		user.Name,
	)

	if err != nil {
		return fmt.Errorf("could not list feeds: %s", err)
	}

	for _, followedfeed := range(followedfeeds) {
		fmt.Println(followedfeed.Feed)
	}

	return nil
}