package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tucond/go-linebot.git/github"
)

func main() {
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// weatherパッケージパッケージから天気情報の文字列をを取得する
	result, err := github.GetText()
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(result)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

//参考：https://qiita.com/yuki_0920/items/cbdbd5220a6a8b4eef19
