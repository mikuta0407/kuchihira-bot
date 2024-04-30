package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type BskyConfig struct {
	Host     string `json:"host"`
	Handle   string `json:"handle"`
	Password string `json:"password"`
	Dir      string
	Prefix   string
}

type TwitterConfig struct {
	APIKey           string `json:"apikey"`
	APIKeySecret     string `json:"apikeysec"`
	OAuthToken       string `json:"oauthtoken"`
	OAuthTokenSecret string `json:"oauthtokensec"`
}

type DiscordConfig struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	URL       string `json:"url"`
}

type KuchihiraConfig struct {
	Hashtag         string   `json:"hashtag"`
	RSSURL          string   `json:"rss"`
	OmnyFMURL       string   `json:"omnyfm"`
	VoicyURL        string   `json:"voicy"`
	TwitterMentions []string `json:"mentions"`
}

func GetConfigDir() (string, error) {
	// nowdir, _ := os.Getwd()
	// return filepath.Join(nowdir, "_config"), nil

	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(exe), "_config"), nil
}

func LoadBskyConfig() (*BskyConfig, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "config-bsky.json")

	os.MkdirAll(filepath.Dir(path), 0700)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}
	var cfg BskyConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}
	if cfg.Host == "" {
		cfg.Host = "https://bsky.social"
	}
	cfg.Dir = dir
	return &cfg, nil
}

func LoadTwitterConfig() (*TwitterConfig, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "config-twtr.json")

	os.MkdirAll(filepath.Dir(path), 0700)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}
	var cfg TwitterConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	return &cfg, nil
}

func LoadDiscordConfig() (*DiscordConfig, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "config-discord.json")

	os.MkdirAll(filepath.Dir(path), 0700)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}
	var cfg DiscordConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	return &cfg, nil
}

func LoadKuchihiraConfig() (*KuchihiraConfig, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "config-kuchihira.json")

	os.MkdirAll(filepath.Dir(path), 0700)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}
	var cfg KuchihiraConfig
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file: %w", err)
	}

	return &cfg, nil
}
