package cli

import (
	"context"
	"fmt"

	"github.com/SirVoly/gator/internal/rssfeed"
)

func handlerAgg(s *state, cmd command) error {
	rf, err := rssfeed.FetchFeed(
		context.Background(),
		"https://www.wagslane.dev/index.xml",
	)

	fmt.Println(rf.SPrint())

	return err
}
