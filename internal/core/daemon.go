package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/mikuta0407/kuchihira-bot/internal/discord"
)

func DaemonStart(isDebug bool) error {
	if isDebug {
		fmt.Println("==== DRY RUN MODE!!! ====")
	}

	// RSS取得
	var items []Item
	var latestGUID string
	var err error
	for {
		items, latestGUID, err = getNewPost(kuchihiraCfg.RSSURL)
		if err != nil && !errors.Is(err, ErrorNoUpdate) {
			// discord.DoPost(discordCfg, "Failed: Get RSS, "+err.Error(), isDebug)
			fmt.Println(err)
		}

		if latestGUID != "" {
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

		}
		time.Sleep(time.Second * 20)

	}
}
