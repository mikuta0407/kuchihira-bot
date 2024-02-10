package twitter

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"github.com/mikuta0407/kuchihira-bot/internal/config"
)

func DoPost(cfg *config.TwitterConfig, text string) error {
	// Setup credentials
	os.Setenv("GOTWI_API_KEY", cfg.APIKey)
	os.Setenv("GOTWI_API_KEY_SECRET", cfg.APIKeySecret)
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           cfg.OAuthToken,
		OAuthTokenSecret:     cfg.OAuthTokenSecret,
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		return err
	}

	p := &types.CreateInput{
		Text: &text,
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return err
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
	return nil
}
