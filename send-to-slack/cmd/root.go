package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "send-to-slack",
	Short: "ISUCON 用 Slack へのメッセージ・ファイル送信ツール",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// グローバルコマンドオプション
var (
	headerTitle string
	noHeader    bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&headerTitle, "header-title", "t", "", "共通のヘッダーメッセージのタイトル")
	rootCmd.PersistentFlags().BoolVarP(&noHeader, "no-header", "", false, "共通のヘッダーメッセージを投稿しない")
	rootCmd.MarkFlagsOneRequired("header-title", "no-header")
	rootCmd.MarkFlagsMutuallyExclusive("header-title", "no-header")
}
