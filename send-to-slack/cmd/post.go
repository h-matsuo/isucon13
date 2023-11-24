/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/h-matsuo/isucon13/send-to-slack/lib"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Slack にメッセージを投稿する (Markdown)",
	Example: `  共通ヘッダーメッセージを付けてメッセージを投稿する:
    send-to-slack post \
      --header-title 'これはヘッダーのタイトルです' \
      --message 'これはメッセージです'

  共通ヘッダーメッセージを省略して標準入力をメッセージとして投稿する:
    send-to-slack upload \
      --no-header \
      --stdin`,
	Run: runPost,
}

// コマンドオプション
var (
	message string
	stdin   bool
)

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&message, "message", "m", "", "投稿するメッセージ")
	postCmd.Flags().BoolVarP(&stdin, "stdin", "s", false, "標準入力を投稿するメッセージとして指定")
	postCmd.MarkFlagsOneRequired("message", "stdin")
	postCmd.MarkFlagsMutuallyExclusive("message", "stdin")
}

func runPost(cmd *cobra.Command, args []string) {
	if stdin {
		fmt.Scan(&message)
	}

	api := lib.GetSlackClient()
	messageBlocks := []slack.Block{}
	if !noHeader {
		messageBlocks = append(messageBlocks, lib.GetHeaderMessageBlocks(headerTitle)...)
	}
	messageBlocks = append(messageBlocks,
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: message,
			},
		},
	)
	_, _, err := api.PostMessage(lib.GetSlackChannelID(), slack.MsgOptionBlocks(messageBlocks...))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Slack へのメッセージの投稿に失敗しました")
		panic(err)
	}
}
