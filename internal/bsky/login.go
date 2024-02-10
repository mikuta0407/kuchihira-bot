package bsky

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikuta0407/kuchihira-bot/internal/config"
)

func DoLogin(host, handle, password string) error {
	dir, err := config.GetConfigDir()
	if err != nil {
		return err
	}
	var cfg config.BskyConfig
	cfg.Host = host
	cfg.Handle = handle
	cfg.Password = password

	b, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot make config file: %w", err)
	}
	err = os.WriteFile(filepath.Join(dir, "config-bsky.json"), b, 0644)
	if err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}
	fmt.Println(filepath.Join(dir, "config-bsky.json"))
	return nil
}
