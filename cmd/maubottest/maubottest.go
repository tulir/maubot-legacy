package main

import (
	"fmt"

	"maunium.net/go/maubot"
	"maunium.net/go/maubot/irc"
	"maunium.net/go/maubot/stdio"
	"maunium.net/go/maubot/telegram"
	"maunium.net/go/mauflag"
)

var tgKey = mauflag.MakeFull("t", "telegram", "The Telegram bot secret to use.", "").String()
var ircServ = mauflag.MakeFull("i", "irc", "The IRC server host:port to connect to", "").String()
var ircUser = mauflag.MakeFull("u", "user", "The IRC user to use", "maubottest").String()
var ircNick = mauflag.MakeFull("n", "nick", "The IRC nick to use", "maubottest").String()
var ircRealname = mauflag.MakeFull("r", "realname", "The IRC realname to use", "A Maubot Test").String()
var useStdIO = mauflag.MakeFull("o", "stdio", "Whether or not to enable the stdio handler", "false").Bool()
var wantHelp, _ = mauflag.MakeHelpFlag()

func main() {
	mauflag.Parse()
	mauflag.SetHelpTitles("maubottest - A simple testing utility for maubot", "maubottest [-i] [-t secret]")
	if *wantHelp {
		mauflag.PrintHelp()
		return
	}
	maubot := maubot.New()

	if len(*tgKey) > 0 {
		tg, err := telegram.New(*tgKey)
		if err != nil {
			panic(err)
		}
		maubot.Add(tg)
		tg.Connect()
	}
	if len(*ircServ) > 0 {
		irk, err := irc.New(*ircNick, *ircUser, *ircRealname, *ircServ)
		if err != nil {
			panic(err)
		}
		maubot.Add(irk)
		irk.Connect()
	}
	if *useStdIO {
		io := stdio.New()
		maubot.Add(io)
		io.Connect()
	}

	for message := range maubot.Messages() {
		if message.Text() == "!leave" {
			message.Reply("3:")
			maubot.Remove(message.Source().UID())
			message.Source().Disconnect()
			return
		}
		message.ReplyWithRef(fmt.Sprintf("You said %s", message.Text()))
	}
}
