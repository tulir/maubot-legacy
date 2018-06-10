package matrix

import (
	"fmt"

	"maunium.net/go/maubot-legacy"
	"maunium.net/go/mautrix"
)

// Message is an implementation of maubot.Message for Matrix messages.
type Message struct {
	internal mautrix.Event
	bot      *Bot
}

// Underlying returns the underlying mautrix event object
func (msg *Message) Underlying() interface{} {
	return msg.internal
}

// Source returns the Matrix Bot parent of this message.
func (msg *Message) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the room the message came from.
func (msg *Message) Reply(message string) {
	msg.bot.sendMessage(msg.internal.Room.ID, message)
}

// ReplyWithRef sends a message to the room the message came from
// with a reference to the original message or the sender.
func (msg *Message) ReplyWithRef(message string) {
	msg.bot.sendMessage(msg.internal.Room.ID, fmt.Sprintf("@%s %s", msg.Sender(), message))
}

// Text returns the text in the message
func (msg *Message) Text() string {
	return msg.internal.Content["body"].(string)
}

// Room returns the display name of the current channel or user.
func (msg *Message) Room() string {
	// FIXME I doubt Room.Name is what the Matrix spec wants clients to always show.
	return msg.internal.Room.Name
}

// RoomID returns the ID of the current channel or user.
func (msg *Message) RoomID() string {
	return msg.internal.Room.ID
}

// SenderID returns the ID of the sender.
func (msg *Message) SenderID() string {
	return msg.internal.Sender
}

// Sender returns the preferred displayname of the sender
func (msg *Message) Sender() string {
	// TODO properly implement this (i.e. get displayname)
	return msg.SenderID()
}
