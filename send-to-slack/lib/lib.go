package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func GetSlackClient() *slack.Client {
	return slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN"))
}

func GetSlackChannelID() string {
	return os.Getenv("SLACK_CHANNEL_ID")
}

func GetHeaderMessageBlocks(title string) []slack.Block {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ホスト名の取得に失敗しました")
		panic(err)
	}

	return []slack.Block{
		&slack.HeaderBlock{
			Type: slack.MBTHeader,
			Text: &slack.TextBlockObject{
				Type: "plain_text",
				Text: os.Getenv("SERVER_ICON") + "｜" + title,
			},
		},

		&slack.ContextBlock{
			Type: slack.MBTContext,
			ContextElements: slack.ContextElements{
				Elements: []slack.MixedElement{
					&slack.TextBlockObject{
						Type: "plain_text",
						Text: time.Now().Format("2006/01/02 15:04:05"),
					},
					&slack.TextBlockObject{
						Type: "plain_text",
						Text: "Unit: " + os.Getenv("SERVER_NAME") + " (" + hostname + ")",
					},
				},
			},
		},

		slack.NewDividerBlock(),
	}
}
