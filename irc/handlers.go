// Package irc provides a maubot implementation for IRC.
package irc

import (
	msg "github.com/sorcix/irc"
)

func (bot *Bot) handlePrivmsg(evt *msg.Message) {
	channel := evt.Params[0]
	if evt.Name == "maubottest" {
		return
	} else if bot.nick == channel {
		channel = evt.Name
	}
	bot.SendToListeners(&Message{
		bot:      bot,
		internal: evt,
		sender:   evt.Name,
		message:  evt.Trailing,
		channel:  channel,
	})
}
