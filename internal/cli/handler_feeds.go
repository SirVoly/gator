package cli

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(
		context.Background(),
	)
	if err != nil {
		return fmt.Errorf("couldn't get all feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("\t%s\n", feed.Url)
		fmt.Printf("\tCreated by: %s\n", feed.User)
	}


	return nil
}

