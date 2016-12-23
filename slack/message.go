package slack

import (
	"github.com/nlopes/slack"
	"maunium.net/go/maubot"
)

type SlackMessage struct {
	internal slack.Msg
	bot      *SlackBot
}

// Underlying returns the underlying Slack message object
func (msg *SlackMessage) Underlying() interface{} {
	return msg.internal
}

// Source returns the Slack parent of this message.
func (msg *SlackMessage) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the room the message came from.
func (msg *SlackMessage) Reply(message string) {
	reply := msg.bot.internal.NewOutgoingMessage(message, msg.internal.Channel)
	msg.bot.internal.SendMessage(reply)
}

// ReplyWithRef sends a message to the room the message came from
// with a reference to the original message or the sender.
func (msg *SlackMessage) ReplyWithRef(message string) {
	reply := msg.bot.internal.NewOutgoingMessage(message, msg.internal.Channel)
	msg.bot.internal.SendMessage(reply)
}

// Text returns the text in the message
func (msg *SlackMessage) Text() string {
	return msg.internal.Text
}

// Room returns the display name of the current channel or user.
func (msg *SlackMessage) Room() string {
	if len(msg.internal.Name) > 0 {
		return msg.internal.Name
	}

	return msg.Sender()
}

// RoomID returns the ID of the current channel or user.
func (msg *SlackMessage) RoomID() string {
	return msg.internal.Channel
}

// SenderID returns the ID of the sender.
func (msg *SlackMessage) SenderID() string {
	return msg.internal.User
}

// Sender returns the preferred displayname of the sender
func (msg *SlackMessage) Sender() string {
	return msg.internal.Username
}
