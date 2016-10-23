package main

import (
	"fmt"
	"maunium.net/go/maubot"
	"maunium.net/go/maubot/telegram"
	"maunium.net/go/mauflag"
)

var tgKey = mauflag.MakeKey("t", "telegram").String()

func main() {
	mauflag.Parse()
	maubot := maubot.Create()

	if len(*tgKey) > 0 {
		tg, err := telegram.New(*tgKey)
		if err != nil {
			panic(err)
		}
		maubot.Add(tg)
		tg.Connect()
	}

	for message := range maubot.Messages {
		fmt.Printf("<%s> %s\n", message.Room(), message.Text())
		if len(message.Text()) > 10 {
			message.Reply("*Yay?!*")
		}
	}
}
