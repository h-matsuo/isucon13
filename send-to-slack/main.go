package main

import (
	"fmt"
	"log"
	"os"

	"github.com/h-matsuo/isucon13/send-to-slack/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: .env ファイルの読み込みに失敗しました")
		panic(err)
	}
	if os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN") == "" {
		log.Fatalln("Error: 環境変数 SLACK_BOT_USER_OAUTH_TOKEN がセットされていません")
	}
	if os.Getenv("SLACK_CHANNEL_ID") == "" {
		log.Fatalln("Error: 環境変数 SLACK_CHANNEL_ID がセットされていません")
	}
	if os.Getenv("SERVER_NAME") == "" {
		log.Fatalln("Error: 環境変数 SERVER_NAME がセットされていません")
	}
	if os.Getenv("SERVER_ICON") == "" {
		log.Fatalln("Error: 環境変数 SERVER_ICON がセットされていません")
	}

	// Cobra 実行
	cmd.Execute()
}
