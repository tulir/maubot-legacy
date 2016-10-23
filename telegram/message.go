// Package telegram provides a maubot implementation for Telegram.
package telegram

import (
	"github.com/tucnak/telebot"
	"maunium.net/go/maubot"
	"strconv"
	"strings"
)

// TGMessage is an implementation of the Maubot message for Telegram messages
type TGMessage struct {
	internal telebot.Message
	bot      *TGBot
}

// Underlying returns the underlying Telebot message object
func (msg *TGMessage) Underlying() interface{} {
	return msg.internal
}

// Source returns the TGBot parent of this message.
func (msg *TGMessage) Source() maubot.Bot {
	return msg.bot
}

// Reply sends a message to the room the message came from.
func (msg *TGMessage) Reply(message string) {
	msg.bot.internal.SendMessage(msg.internal.Chat, message, nil)
}

// ReplyWithRef sends a message to the room the message came from
// with a reference to the original message or the sender.
func (msg *TGMessage) ReplyWithRef(message string) {
	// TODO use Telegram reply functionality
	msg.bot.internal.SendMessage(msg.internal.Chat, message, nil)
}

// Text returns the text in the message
func (msg *TGMessage) Text() string {
	return msg.internal.Text
}

// Room returns the display name of the current channel or user.
func (msg *TGMessage) Room() string {
	if len(msg.internal.Chat.Title) > 0 {
		return msg.internal.Chat.Title
	}

	return msg.Sender()
}

// RoomID returns the ID of the current channel or user.
func (msg *TGMessage) RoomID() string {
	return strconv.FormatInt(msg.internal.Chat.ID, 10)
}

// SenderID returns the ID of the sender.
func (msg *TGMessage) SenderID() string {
	return strconv.Itoa(msg.internal.Sender.ID)
}

// Sender returns the preferred displayname of the sender
func (msg *TGMessage) Sender() string {
	fullName := strings.TrimSpace(msg.internal.Sender.FirstName + " " + msg.internal.Sender.LastName)
	if len(fullName) > 0 {
		return fullName
	}

	if len(msg.internal.Sender.Username) > 0 {
		return msg.internal.Sender.Username
	}

	return strconv.Itoa(msg.internal.Sender.ID)
}
