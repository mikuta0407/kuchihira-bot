package rss

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title   string
	URL     string
	PubDate time.Time
}

var jst *time.Location

func init() {
	jst, _ = time.LoadLocation("Asia/Tokyo")
}

func GetLatestRssPost(url string) (Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return Item{}, err
	}

	latestItem := feed.Items[0]
	return Item{
		Title:   latestItem.Title,
		URL:     latestItem.Link,
		PubDate: latestItem.PublishedParsed.In(jst),
	}, nil
}
