package rss

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title string
	URL   string
}

func GetRss(url string) (Item, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fp := gofeed.NewParser()

	nowTime := time.Now().In(jst)
	var pubTime time.Time
	var latestItem *gofeed.Item
	i := 1
	for {
		fmt.Println(i, "回目...")
		feed, err := fp.ParseURL(url)
		latestItem = feed.Items[0]
		if err != nil {
			return Item{}, err
		}
		pubTime = latestItem.PublishedParsed.In(jst)

		if truncateTime(pubTime).Equal(truncateTime(nowTime)) {
			fmt.Println("取得!")
			break
		}

		if i == 5 {
			fmt.Println("失敗しました")
			return Item{}, fmt.Errorf("更新されませんでした")
		}

		time.Sleep(time.Second * 3)
		i++

	}

	return Item{
		Title: latestItem.Title,
		URL:   latestItem.Link,
	}, nil
}

func truncateTime(t time.Time) time.Time {
	t = t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
	return t
}
