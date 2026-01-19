package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <feedName> <feedURL>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	id := uuid.New()
	t := time.Now()

	user_name := s.cfg.GetUser()
	user, err := s.db.GetUser(context.Background(), user_name)
	
	if err != nil {
		return fmt.Errorf("user %s was not found", user_name)
	}

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