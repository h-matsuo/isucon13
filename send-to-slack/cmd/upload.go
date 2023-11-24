package cmd

import (
	"fmt"
	"os"

	"github.com/h-matsuo/isucon13/send-to-slack/lib"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Slack にファイルをアップロード＆チャンネルに投稿する",
	Example: `  共通ヘッダーメッセージを付けてファイルを投稿する:
    send-to-slack upload \
      --header-title 'これはヘッダーのタイトルです' \
      --filepath /path/to/file

  共通ヘッダーメッセージを省略してファイルを投稿する:
    send-to-slack upload \
      --no-header \
      --filepath /path/to/file`,
	Run: runUpload,
}

// コマンドオプション
var (
	filepath string
)

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "アップロードするファイルのパス")
	uploadCmd.MarkFlagRequired("filepath")
}

func runUpload(cmd *cobra.Command, args []string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ファイルの読み込みに失敗しました")
		panic(err)
	}

	api := lib.GetSlackClient()
	slackFile, err := api.UploadFile(
		slack.FileUploadParameters{
			Reader:   file,
			Filename: "Uploaded by send-to-slack",
			Title:    file.Name(),
			// ヘッダーメッセージと一緒にファイルを表示させたいため、チャンネルへの直接投稿はここでは実施しない
			// Channels: []string{lib.GetSlackChannelID()},
		})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Slack へのファイルのアプロードに失敗しました")
		panic(err)
	}

	messageBlocks := []slack.Block{}
	if !noHeader {
		messageBlocks = append(messageBlocks, lib.GetHeaderMessageBlocks(headerTitle)...)
	}
	messageBlocks = append(messageBlocks,
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: ":file_folder: " + slackFile.Permalink,
			},
		},
	)
	_, _, err = api.PostMessage(lib.GetSlackChannelID(), slack.MsgOptionBlocks(messageBlocks...))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Slack へのメッセージの投稿に失敗しました")
		panic(err)
	}
}
