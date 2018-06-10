// Package irc provides a maubot implementation for IRC.
package irc

import (
	msg "github.com/sorcix/irc"
	"maunium.net/go/maubot-legacy"
)

// Message is an implementation of maubot.Message for IRC messages
type Message struct {
	internal *msg.Message
	channel  string
	sender   string
	message  string
	bot      *Bot
}

// Underlying returns the underlying IRC message object
func (msg *Message) Underlying() interface{} {
	return msg.internal
}

// Source returns the IRC Bot parent of this message.
func (msg *Message) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the channel the message came from.
func (msg *Message) Reply(message string) {
	msg.bot.internal.Privmsg(msg.channel, message)
}

// ReplyWithRef sends a message to the channel the message came from
// with a highlight for the sender of the original message.
func (msg *Message) ReplyWithRef(message string) {
	msg.bot.internal.Privmsg(msg.channel, msg.sender+": "+message)
}

// Text returns the text in the message
func (msg *Message) Text() string {
	return msg.message
}

// Room returns the name of the channel.
func (msg *Message) Room() string {
	return msg.channel
}

// RoomID returns the name of the channel.
func (msg *Message) RoomID() string {
	return msg.channel
}

// SenderID returns the nick of the sender.
func (msg *Message) SenderID() string {
	return msg.sender
}

// Sender returns the nick of the sender.
func (msg *Message) Sender() string {
	return msg.sender
}
