package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// トークンを格納
const token = ""
// BotIDを文字列型として宣言
var BotID string

func main() {

	// discordのセッションを作成
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// ユーザーIDまたは現在のユーザーIDのショートカットである「@me」
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// ユーザーidをBotIDに格納
	BotID = u.ID

	// 音声/スピーキング/アップデートのイベントハンドラを追加
	dg.AddHandler(messageHandler)

	// DiscordへのWebSocket接続を作成
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遅延実行：セッションを閉じる
	defer dg.Close()

	fmt.Println("ボットは稼働中です", time.Now())
	<-make(chan bool) // プログラムが自動的に閉じないように
	return
}

// *discordgo.Session は DiscordAPIへの接続を表す。
// *discordgo.MessageCreate は メッセージ作成のイベントデータ（型）
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// メッセージの作成者(この場合はユーザーID) と BotID が等しいなら終了
	if m.Author.ID == BotID {
		return
	}

	// m.Content は discordのメッセージの内容
	// m.COntent に 自由な指定の文字列を。
	if m.Content == "" {
		
		// 指定されたチャネルにメッセージを送信。
		// m.ChannelID：チャネルのID。 m.Content：BOTが送信するメッセージ。
		_, _ = s.ChannelMessageSend(m.ChannelID, "やあみんな！僕は困ったロボ！みんなをサポートする役目がある！")
	}
}