package core

import (
	"fmt"

	"github.com/mikuta0407/kuchihira-bot/internal/bsky"
	"github.com/mikuta0407/kuchihira-bot/internal/config"
	"github.com/mikuta0407/kuchihira-bot/internal/discord"
	"github.com/mikuta0407/kuchihira-bot/internal/rss"
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
func Start() error {

	// RSS取得
	item, err := rss.GetRss(kuchihiraCfg.RSSURL)
	if err != nil {
		discord.DoPost(discordCfg, "Failed: Get RSS, "+err.Error())
		return err
	}

	// Twitterへの投稿
	// 本文生成
	text, err := generateTwitterPostText(item, kuchihiraCfg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(text)
	// 投稿
	// if err := twitter.DoPost(twtrCfg, text); err != nil {
	// 	discord.DoPost(discordCfg, "Failed: Twitter")
	// }

	// Blueskyへの投稿
	// 本文生成
	text, err = generateBlueskyPostText(item, kuchihiraCfg)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	fmt.Println(text)
	if err := bsky.DoPost(bskyCfg, text); err != nil {
		discord.DoPost(discordCfg, "Failed: Bluesky")
	}

	discord.DoPost(discordCfg, "Kuchihira-bot Finish")
	return nil
}
