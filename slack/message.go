package slack

import (
	"github.com/nlopes/slack"
	"maunium.net/go/maubot"
)

type Message struct {
	internal slack.Msg
	bot      *Bot
}

// Underlying returns the underlying Slack message object
func (msg *Message) Underlying() interface{} {
	return msg.internal
}

// Source returns the Slack parent of this message.
func (msg *Message) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the room the message came from.
func (msg *Message) Reply(message string) {
	reply := msg.bot.internal.NewOutgoingMessage(message, msg.internal.Channel)
	msg.bot.internal.SendMessage(reply)
}

// ReplyWithRef sends a message to the room the message came from
// with a reference to the original message or the sender.
func (msg *Message) ReplyWithRef(message string) {
	reply := msg.bot.internal.NewOutgoingMessage(message, msg.internal.Channel)
	msg.bot.internal.SendMessage(reply)
}

// Text returns the text in the message
func (msg *Message) Text() string {
	return msg.internal.Text
}

// Room returns the display name of the current channel or user.
func (msg *Message) Room() string {
	if len(msg.internal.Name) > 0 {
		return msg.internal.Name
	}

	return msg.Sender()
}

// RoomID returns the ID of the current channel or user.
func (msg *Message) RoomID() string {
	return msg.internal.Channel
}

// SenderID returns the ID of the sender.
func (msg *Message) SenderID() string {
	return msg.internal.User
}

// Sender returns the preferred displayname of the sender
func (msg *Message) Sender() string {
	return msg.internal.Username
}
