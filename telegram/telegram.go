// Package telegram provides a maubot implementation for Telegram.
package telegram

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/tucnak/telebot"
	"maunium.net/go/maubot-legacy"
)

// SimpleRecipient is an implementation of the telebot Recipient interface.
type SimpleRecipient struct {
	Recipient string
}

// Destination returns the Recipient field.
func (sr SimpleRecipient) Destination() string {
	return sr.Recipient
}

// New creates an instance of the maubot.Bot implementation for Telegram.
func New(token string) (maubot.Bot, error) {
	bot := &Bot{internal: nil, token: token, uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}, stop: make(chan bool, 1)}
	return bot, nil
}

// Bot is an implementation of maubot.Bot for Telegram.
type Bot struct {
	internal  *telebot.Bot
	listeners []chan maubot.Message
	uid       string
	token     string
	connected bool
	stop      chan bool
}

// Connect connects to the Telegram servers.
func (bot *Bot) Connect() error {
	tg, err := telebot.NewBot(bot.token)
	if err != nil {
		return err
	}
	bot.internal = tg

	go func() {
		messages := make(chan telebot.Message)
		tg.Listen(messages, 1*time.Second)
		bot.connected = true
		for {
			select {
			case message := <-messages:
				bot.SendToListeners(&Message{bot: bot, internal: message})
			case stop := <-bot.stop:
				if stop {
					bot.connected = false
					return
				}
			}
		}
	}()
	return nil
}

// UID returns the unique ID for this instance.
func (bot *Bot) UID() string {
	return bot.uid
}

// Connected returns whether or not the message listener is active.
func (bot *Bot) Connected() bool {
	return bot.connected
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *Bot) Disconnect() error {
	bot.stop <- true
	return nil
}

// Underlying returns the telebot bot object.
func (bot *Bot) Underlying() interface{} {
	return bot.internal
}

// SendMessage sends a message to the given channel or user.
func (bot *Bot) SendMessage(msg maubot.OutgoingMessage) {
	bot.internal.SendMessage(SimpleRecipient{Recipient: msg.RoomID}, msg.Text, nil)
}

// SendToListeners sends the given message to all listener channels.
func (bot *Bot) SendToListeners(message maubot.Message) {
	for _, listener := range bot.listeners {
		listener <- message
	}
}

// AddListener adds a message listener
func (bot *Bot) AddListener(listener chan maubot.Message) {
	bot.listeners = append(bot.listeners, listener)
}
