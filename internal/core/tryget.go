package core

import (
	"fmt"
	"time"

	"github.com/mikuta0407/kuchihira-bot/internal/rss"
)

func getNewPost(url string) (item rss.Item, err error) {

	nowTime := time.Now().In(jst)
	var i int
	for i = 1; i <= 360; i++ {
		fmt.Println(i, "回目...")
		item, err = rss.GetLatestRssPost(url)
		if err != nil {
			return rss.Item{}, err
		}

		if truncateTime(item.PubDate).Equal(truncateTime(nowTime)) {
			fmt.Println("取得!")
			break
		}

		time.Sleep(time.Second * 20)
	}
	if i == 361 {
		fmt.Println("失敗しました")
		return rss.Item{}, fmt.Errorf("更新されませんでした")
	}

	return

}

func truncateTime(t time.Time) time.Time {
	t = t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
	return t
}
