package main

import (
	"log"
	"os"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// ハンドラの登録
	http.HandleFunc("/callback", callbackHandler)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// BOTを初期化
    bot, err := linebot.New(
        os.Getenv("LINE_BOT_CHANNEL_SECRET"),
        os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
    )
    // エラーに値があればログに出力し終了する
    if err != nil {
        log.Fatal(err)
    }

	// リクエストからBOTのイベントを取得
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			// メッセージが位置情報の場合
			case *linebot.LocationMessage:
//				sendRestoInfo(bot, event)
			}
			// 他にもスタンプや画像、位置情報など色々受信可能
		}
	}
}
