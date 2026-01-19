package rssfeed

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return nil, fmt.Errorf("Error in http request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Error in http response: %w", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	rf := RSSFeed{}
	xml.Unmarshal(body, &rf)

	rf.Unescape()

	return &rf, nil
}
