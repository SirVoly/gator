package cli

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	followedfeeds, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		s.cfg.CurrentUserName,
	)

	if err != nil {
		return fmt.Errorf("could not list feeds: %s", err)
	}

	for _, followedfeed := range(followedfeeds) {
		fmt.Println(followedfeed.Feed)
	}

	return nil
}