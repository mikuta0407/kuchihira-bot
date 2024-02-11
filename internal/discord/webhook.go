package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikuta0407/kuchihira-bot/internal/config"
)

type Discord struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}

func DoPost(cfg *config.DiscordConfig, text string, isDebug bool) error {

	if isDebug {
		fmt.Println("========== Discord ==========")
		fmt.Println(text)
		return nil
	}

	discord := Discord{
		Username:  cfg.Username,
		AvatarUrl: cfg.AvatarUrl,
		Content:   text,
	}

	// encode json
	discord_json, _ := json.Marshal(discord)
	fmt.Println(string(discord_json))

	// discord webhook_url
	webhook_url := cfg.URL
	res, err := http.Post(
		webhook_url,
		"application/json",
		bytes.NewBuffer(discord_json),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
