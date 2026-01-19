package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <feedName> <feedURL>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	id := uuid.New()
	t := time.Now()

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        id,
			CreatedAt: t,
			UpdatedAt: t,
			Name:      name,
			Url:	   url,
			UserID:	   user.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("feed %s already exists", url)
	}

	follow_id := uuid.New()
	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:			follow_id,
			CreatedAt: 	t,
			UpdatedAt: 	t,
			UserID: 	user.ID,
			FeedID: 	feed.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("could not add user as a follower of the feed")
	}
	
	fmt.Printf("Feed was created: \n")

	// User Log
	fmt.Printf("Feed: %s\n", feed.Name)
	fmt.Printf("\tid: %s\n", feed.ID)
	fmt.Printf("\tcreated_at: %s\n", feed.CreatedAt)
	fmt.Printf("\tupdated_at: %s\n", feed.UpdatedAt)
	fmt.Printf("\tURL: %s\n", feed.Url)
	fmt.Printf("\tUser ID: %s\n", feed.UserID)

	return nil
}