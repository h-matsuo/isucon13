# send-to-slack

ISUCON 競技中に、サーバーから Slack にメッセージやファイルを投稿できるツール。

このツールは、リポジトリルートにある [Taskfile](https://taskfile.dev/) 経由で、  
インストールせずに `go run main.go` で直接実行することを想定しています。

# Setup

```sh
# サーバーを Git 管理下に置いた状態で実行
$ cd /path/to/send-to-slack
$ go mod download

# .env ファイルを準備
$ cp .env.template .env
$ vim .env
```
