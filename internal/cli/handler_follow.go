package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	
	id := uuid.New()
	t := time.Now()

	user, err := s.db.GetUser(context.Background(), s.cfg.GetUser())
	if err != nil {
		return fmt.Errorf("user %s was not found", s.cfg.GetUser())
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("feed %s was not found", url)
	}

	feed_follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:			id,
			CreatedAt: 	t,
			UpdatedAt: 	t,
			UserID: 	user.ID,
			FeedID: 	feed.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("user already follows the feed")
	}

	fmt.Printf("Follow was created: \n")

	// User Log
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("\tFeed: %s\n", feed_follow.Feed)
	fmt.Printf("\tUser: %s\n", feed_follow.User)

	return nil
}