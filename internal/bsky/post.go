package bsky

import (
	"context"
	"fmt"
	"time"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/mikuta0407/kuchihira-bot/internal/config"
)

func DoPost(cfg *config.BskyConfig, text string, isDebug bool) error {

	if isDebug {
		fmt.Println("========== Bluesky ==========")
		fmt.Println(text)
		return nil
	}

	xrpcc, err := makeXRPCC(cfg)
	if err != nil {
		return fmt.Errorf("cannot create client: %w", err)
	}

	post := &bsky.FeedPost{
		Text:      text,
		CreatedAt: time.Now().Local().Format(time.RFC3339),
	}

	linkentries := extractLinksBytes(text)
	for i, entry := range linkentries {
		post.Facets = append(post.Facets, &bsky.RichtextFacet{
			Features: []*bsky.RichtextFacet_Features_Elem{
				{
					RichtextFacet_Link: &bsky.RichtextFacet_Link{
						Uri: entry.text,
					},
				},
			},
			Index: &bsky.RichtextFacet_ByteSlice{
				ByteStart: entry.start,
				ByteEnd:   entry.end,
			},
		})

		if post.Embed == nil {
			post.Embed = &bsky.FeedPost_Embed{}
		}

		// 最後のものだけカード化する
		if len(linkentries)-1 == i {
			if post.Embed == nil {
				post.Embed = &bsky.FeedPost_Embed{}
			}
			if post.Embed.EmbedExternal == nil {
				addLink(xrpcc, post, entry.text)
			}
		}
	}

	resp, err := comatproto.RepoCreateRecord(context.TODO(), xrpcc, &comatproto.RepoCreateRecord_Input{
		Collection: "app.bsky.feed.post",
		Repo:       xrpcc.Auth.Did,
		Record: &lexutil.LexiconTypeDecoder{
			Val: post,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}
	fmt.Println(resp.Uri)

	return nil
}
