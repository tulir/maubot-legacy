// Package telegram provides a maubot implementation for Telegram.
package telegram

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tucnak/telebot"
	"maunium.net/go/maubot"
)

var markdown *telebot.SendOptions

func init() {
	markdown = new(telebot.SendOptions)
	markdown.ParseMode = telebot.ModeMarkdown
}

// Message is an implementation of maubot.Message for Telebot messages
type Message struct {
	internal telebot.Message
	bot      *Bot
}

// Underlying returns the underlying Telebot message object
func (msg *Message) Underlying() interface{} {
	return msg.internal
}

// Source returns the TGBot parent of this message.
func (msg *Message) Source() maubot.Bot {
	return msg.bot
}

// FormatMessage formats a Markdown message into Telegram format
func FormatMessage(message string) string {
	message = strings.Replace(message, "**", "*", -1)
	return message
}

// Reply sends a message to the room the message came from.
func (msg *Message) Reply(message string) {
	msg.bot.internal.SendMessage(msg.internal.Chat, FormatMessage(message), markdown)
}

// ReplyWithRef sends a message to the room the message came from
// with a reference to the original message or the sender.
func (msg *Message) ReplyWithRef(message string) {
	ref := strconv.Itoa(msg.internal.Sender.ID)
	if len(msg.internal.Sender.Username) > 0 {
		ref = msg.internal.Sender.Username
	}
	msg.bot.internal.SendMessage(msg.internal.Chat, fmt.Sprintf("@%s %s", ref, FormatMessage(message)), markdown)
}

// Text returns the text in the message
func (msg *Message) Text() string {
	return msg.internal.Text
}

// Room returns the display name of the current channel or user.
func (msg *Message) Room() string {
	if len(msg.internal.Chat.Title) > 0 {
		return msg.internal.Chat.Title
	}

	return msg.Sender()
}

// RoomID returns the ID of the current channel or user.
func (msg *Message) RoomID() string {
	return strconv.FormatInt(msg.internal.Chat.ID, 10)
}

// SenderID returns the ID of the sender.
func (msg *Message) SenderID() string {
	return strconv.Itoa(msg.internal.Sender.ID)
}

// Sender returns the preferred displayname of the sender
func (msg *Message) Sender() string {
	fullName := strings.TrimSpace(msg.internal.Sender.FirstName + " " + msg.internal.Sender.LastName)
	if len(fullName) > 0 {
		return fullName
	}

	if len(msg.internal.Sender.Username) > 0 {
		return msg.internal.Sender.Username
	}

	return strconv.Itoa(msg.internal.Sender.ID)
}
