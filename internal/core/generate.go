package core

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/mikuta0407/kuchihira-bot/internal/config"
)

//go:embed templates
var templates embed.FS

type PostVars struct {
	Title     string
	Hashtag   string
	OmnyFMURL string
	VoicyURL  string
	Mentions  []string
	ItemURL   string
}

func generateTwitterPostText(item Item, kuchihiraCfg *config.KuchihiraConfig) (string, error) {
	postVars := PostVars{
		Title:     item.Title,
		Hashtag:   kuchihiraCfg.Hashtag,
		OmnyFMURL: kuchihiraCfg.OmnyFMURL,
		VoicyURL:  kuchihiraCfg.VoicyURL,
		Mentions:  kuchihiraCfg.TwitterMentions,
		ItemURL:   item.URL,
	}

	body, err := generateBodyFromTemplate("twitter.tmpl", &postVars)
	if err != nil {
		return "", err
	}

	return body, nil
}

func generateBlueskyPostText(item Item, kuchihiraCfg *config.KuchihiraConfig, isDebug bool) (string, error) {
	postVars := PostVars{
		Title:     item.Title,
		Hashtag:   kuchihiraCfg.Hashtag,
		OmnyFMURL: kuchihiraCfg.OmnyFMURL,
		VoicyURL:  kuchihiraCfg.VoicyURL,
		ItemURL:   item.URL,
	}

	body, err := generateBodyFromTemplate("bluesky.tmpl", &postVars)
	if err != nil {
		return "", err
	}

	return body, nil
}

func generateBodyFromTemplate(tName string, value any) (string, error) {
	t, err := template.ParseFS(templates, fmt.Sprintf("templates/%s", tName))
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, value); err != nil {
		return "", err
	}
	return buf.String(), nil
}
