package rssfeed

import (
	"fmt"
	"html"
	"strings"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (rf RSSFeed) SPrint() string {
	var result []string

	result = append(result, fmt.Sprintf("Channel: %s", rf.Channel.Title))
	result = append(result, "-------")
	result = append(result, fmt.Sprintf("Link: %s", rf.Channel.Link))
	result = append(result, fmt.Sprintf("Description: %s", rf.Channel.Description))
	result = append(result, fmt.Sprintf("Items: %d", len(rf.Channel.Item)))
	for _, item := range rf.Channel.Item {
		result = append(result, "\t- " + strings.Join(item.SPrint(), "\n\t"))
	}

	return strings.Join(result, "\n")
}

func (ri RSSItem) SPrint() []string {
	var result []string

	result = append(result, fmt.Sprintf("%s (%s)", ri.Title, ri.PubDate))
	result = append(result, fmt.Sprintf("\tLink: %s", ri.Link))
	result = append(result, fmt.Sprintf("\tDescription: %s", ri.Description))

	return result
}

func (rf RSSFeed) Unescape() {
	rf.Channel.Title = html.UnescapeString(rf.Channel.Title)
	rf.Channel.Description = html.UnescapeString(rf.Channel.Description)
	for idx, item := range rf.Channel.Item {
		rf.Channel.Item[idx].Title = html.UnescapeString(item.Title)
		rf.Channel.Item[idx].Description = html.UnescapeString(item.Description)
	}
}
