// Package irc provides a maubot implementation for IRC.
package irc

import (
	msg "github.com/sorcix/irc"
	"maunium.net/go/maubot"
)

// IRCMessage is an implementation of the Maubot message for IRC messages
type IRCMessage struct {
	internal *msg.Message
	channel  string
	sender   string
	message  string
	bot      *IRCBot
}

// Underlying returns the underlying IRC message object
func (msg *IRCMessage) Underlying() interface{} {
	return msg.internal
}

// Source returns the IRCBot parent of this message.
func (msg *IRCMessage) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the channel the message came from.
func (msg *IRCMessage) Reply(message string) {
	msg.bot.internal.Privmsg(msg.channel, message)
}

// ReplyWithRef sends a message to the channel the message came from
// with a highlight for the sender of the original message.
func (msg *IRCMessage) ReplyWithRef(message string) {
	msg.bot.internal.Privmsg(msg.channel, msg.sender+": "+message)
}

// Text returns the text in the message
func (msg *IRCMessage) Text() string {
	return msg.message
}

// Room returns the name of the channel.
func (msg *IRCMessage) Room() string {
	return msg.channel
}

// RoomID returns the name of the channel.
func (msg *IRCMessage) RoomID() string {
	return msg.channel
}

// SenderID returns the nick of the sender.
func (msg *IRCMessage) SenderID() string {
	return msg.sender
}

// Sender returns the nick of the sender.
func (msg *IRCMessage) Sender() string {
	return msg.sender
}
