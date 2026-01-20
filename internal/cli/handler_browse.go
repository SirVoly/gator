package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SirVoly/gator/internal/database"
)



func handlerBrowse(s *state, cmd command, u database.User) error {
	var limit int32
	limit = 10
	if len(cmd.Args) != 0 {
		arg, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("usage: %s <amount=10>", cmd.Name)
		}
		limit = int32(arg)
	}

	posts, err := s.db.ListPostsForUser(
		context.Background(),
		database.ListPostsForUserParams{
			UserID: u.ID,
			Limit: limit,
		},
	)

	if err != nil {
		return err
	}
	for _, p := range posts {
		printPost(p)
	}
	return nil
}

func printPost(p database.ListPostsForUserRow) {
	fmt.Printf("# %s\n", p.Title)
	if p.PublishedAt.Valid {
		fmt.Printf("\t%s\n", p.PublishedAt.Time)
	}
	fmt.Printf("\t%s\n", p.Url)
	fmt.Printf("\tFrom %s\n", p.FeedTitle)
	if p.Description.Valid {
		fmt.Printf("## Description\n")
		fmt.Printf("%s\n", p.Description.String)
	}
	fmt.Println()
}