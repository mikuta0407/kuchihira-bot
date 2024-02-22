package core

import (
	"fmt"
	"time"

	"github.com/mikuta0407/kuchihira-bot/internal/bsky"
	"github.com/mikuta0407/kuchihira-bot/internal/config"
	"github.com/mikuta0407/kuchihira-bot/internal/discord"
	"github.com/mikuta0407/kuchihira-bot/internal/twitter"
)

var (
	discordCfg   *config.DiscordConfig
	kuchihiraCfg *config.KuchihiraConfig
	twtrCfg      *config.TwitterConfig
	bskyCfg      *config.BskyConfig
)

func init() {
	var err error
	// Webhook
	discordCfg, err = config.LoadDiscordConfig()
	if err != nil {
		fmt.Println(err)
	}
	// RSSやくちをひらく用
	kuchihiraCfg, err = config.LoadKuchihiraConfig()
	if err != nil {
		fmt.Println(err)
	}
	// Twitter
	twtrCfg, err = config.LoadTwitterConfig()
	if err != nil {
		fmt.Println(err)
	}
	// Bluesky
	bskyCfg, err = config.LoadBskyConfig()
	if err != nil {
		fmt.Println(err)
	}

}
func Start(isDebug bool) error {

	if isDebug {
		fmt.Println("==== DRY RUN MODE!!! ====")
	}

	// RSS取得
	items, latestGUID, err := getNewPost(kuchihiraCfg.RSSURL)
	if err != nil {
		discord.DoPost(discordCfg, "Failed: Get RSS, "+err.Error(), isDebug)
		return err
	}

	if len(items) != 1 {
		discord.DoPost(discordCfg, "Warning: Multi publish detected", isDebug)
	}

	if err := saveLastGUID(latestGUID); err != nil {
		discord.DoPost(discordCfg, "Failed: saveLastGUID: "+latestGUID, isDebug)
	}

	for _, v := range items {
		if err := post(v, isDebug); err != nil {
			return err
		}
		if len(items) != 1 {
			time.Sleep(time.Second * 2)
		}
	}

	return nil
}

func post(item Item, isDebug bool) error {
	// Twitterへの投稿
	// 本文生成
	text, err := generateTwitterPostText(item, kuchihiraCfg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 投稿
	if err := twitter.DoPost(twtrCfg, text, isDebug); err != nil {
		discord.DoPost(discordCfg, "Failed: Twitter", isDebug)
	}

	// Blueskyへの投稿
	// 本文生成
	text, err = generateBlueskyPostText(item, kuchihiraCfg, isDebug)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := bsky.DoPost(bskyCfg, text, isDebug); err != nil {
		discord.DoPost(discordCfg, "Failed: Bluesky", isDebug)
	}

	discord.DoPost(discordCfg, "Kuchihira-bot Finish", isDebug)

	return nil

}
