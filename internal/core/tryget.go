package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikuta0407/kuchihira-bot/internal/rss"
	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title string
	URL   string
}

func getNewPost(url string) (items []Item, latestGUID string, err error) {

	lastGUID, err := loadLastGUID()
	if err != nil {
		return
	}
	var feedItems []*gofeed.Item

	feedItems, err = rss.GetAllRssFeed(url)
	if err != nil {
		return
	}

	if !isExistNewPost(feedItems[0], lastGUID) {
		return items, "", ErrorNoUpdate
	}

	items, latestGUID = extractNewPosts(feedItems, lastGUID)
	return
}

func isExistNewPost(item *gofeed.Item, lastGUID string) bool {
	return item.GUID != lastGUID
}

func extractNewPosts(feeditems []*gofeed.Item, lastGUID string) (newItems []Item, latestGUID string) {
	latestGUID = feeditems[0].GUID

	for _, v := range feeditems {
		if v.GUID == lastGUID {
			return
		}

		tmpItems := []Item{
			{
				Title: v.Title,
				URL:   v.Link,
			},
		}

		newItems = append(tmpItems, newItems...)
		if lastGUID == "" {
			return
		}
	}

	return
}

func getDataDir() (string, error) {
	// nowdir, err := os.Getwd()
	// return filepath.Join(nowdir, "_data"), err

	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(exe), "_data"), nil
}

func loadLastGUID() (string, error) {
	dir, err := getDataDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, "last-guid.json")
	_, err = os.Stat(path)
	if err != nil {
		// なかったので空で返す
		return "", nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot load last guid: %w", err)
	}
	var guidJSON GUIDJSON
	err = json.Unmarshal(b, &guidJSON)
	if err != nil {
		return "", fmt.Errorf("cannot load last guid: %w", err)
	}

	return guidJSON.GUID, nil
}

func saveLastGUID(guid string) error {
	dir, err := getDataDir()
	if err != nil {
		fmt.Println(err)
		return err
	}
	var guidJSON GUIDJSON
	guidJSON.GUID = guid

	b, err := json.Marshal(&guidJSON)
	if err != nil {
		return fmt.Errorf("cannot make guid json: %w", err)
	}
	err = os.WriteFile(filepath.Join(dir, "last-guid.json"), b, 0644)
	if err != nil {
		return fmt.Errorf("cannot write guid file: %w", err)
	}
	return nil
}

type GUIDJSON struct {
	GUID string `json:"guid"`
}
