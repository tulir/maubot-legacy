package main

import (
	"fmt"
	"maunium.net/go/maubot"
	"maunium.net/go/maubot/stdio"
	"maunium.net/go/maubot/telegram"
	"maunium.net/go/mauflag"
)

var tgKey = mauflag.MakeFull("t", "telegram", "The Telegram bot secret to use.", "").String()
var useStdIO = mauflag.MakeFull("i", "stdio", "Whether or not to enable the stdio handler", "false").Bool()
var wantHelp, _ = mauflag.MakeHelpFlag()

func main() {
	mauflag.Parse()
	mauflag.SetHelpTitles("maubottest - A simple testing utility for maubot", "maubottest [-i] [-t secret]")
	if *wantHelp {
		mauflag.PrintHelp()
		return
	}
	maubot := maubot.Create()

	if len(*tgKey) > 0 {
		tg, err := telegram.New(*tgKey)
		if err != nil {
			panic(err)
		}
		maubot.Add(tg)
		tg.Connect()
	}
	if *useStdIO {
		io := stdio.New()
		maubot.Add(io)
		io.Connect()
	}

	for message := range maubot.Messages {
		message.Reply(fmt.Sprintf("Hello, %s. You said %s", message.Sender(), message.Text()))
	}
}
