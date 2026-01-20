package cli

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SirVoly/gator/internal/database"
	"github.com/SirVoly/gator/internal/rssfeed"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		err = addPost(s, item, feed.ID)
		if err != nil {
			// Unexpected error: Stop scraping
			fmt.Println(err)
			return
		}
	}
}

func addPost(s * state, item rssfeed.RSSItem, feed_id uuid.UUID) error {
	id := uuid.New()
	t := time.Now()
	desc := sql.NullString{
		String: item.Description,
		Valid: (item.Description != ""),
	}

	pubDate := sql.NullTime{
		Time: time.Time{},
		Valid: item.PubDate != "",
	}
	if pubDate.Valid {
		parsedTime, err := parseFlexibleTime(item.PubDate)
		if err != nil {
			return fmt.Errorf("could not parse %q: %w", s, err)
		}
		pubDate.Time = parsedTime
	}

	_, err := s.db.CreatePost(
		context.Background(),
		database.CreatePostParams{
			ID: id,
			CreatedAt: t,
			UpdatedAt: t,
			Title: item.Title,
			Url: item.Link,
			Description: desc,
			PublishedAt: pubDate,
			FeedID: feed_id,
		},
	)

	if err != nil {
		pg_err, ok := err.(*pq.Error)
		if ok && pg_err.Code.Name() == "unique_violation" { //Error Code 23505
			return nil
		}
		// Unexpected error
		return err
	}

	// Post is successfully made
	return nil
}

// Help functions

var layouts = []string{
	time.RFC3339,            // "2006-01-02T15:04:05Z07:00"
	"2006-01-02 15:04:05",   // "2024-01-02 15:04:05"
	"2006-01-02",            // "2024-01-02"
	"02/01/2006 15:04",      // "02/01/2024 15:04"
	time.RFC1123Z, // "Tue, 20 Jan 2026 11:15:33 +0000"
}

func parseFlexibleTime(s string) (time.Time, error) {
	var lastErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, fmt.Errorf("could not parse %q: %w", s, lastErr)
}
