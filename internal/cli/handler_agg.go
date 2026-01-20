package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/SirVoly/gator/internal/rssfeed"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_requests>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	if time_between_reqs < time.Duration(5 * float64(time.Second)) {
		return fmt.Errorf("In between time cannot be shorter than 5 seconds.")
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	t := time.Now()

	feed, err = s.db.MarkFeedFetched(
		context.Background(),
		database.MarkFeedFetchedParams{
			UpdatedAt: t,
			ID: feed.ID,
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	rf, err := rssfeed.FetchFeed(
		context.Background(),
		feed.Url,
	)

	for _, item := range(rf.Channel.Item) {
		fmt.Println(item.Title)
	}
}
