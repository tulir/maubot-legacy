// maubot - A chatbot platform abstraction library
// Copyright (C) 2016 Tulir Asokan

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package telegram provides a maubot implementation for Telegram.
package telegram

import (
	"github.com/tucnak/telebot"
	"maunium.net/go/maubot"
	"time"
)

// New creates an instance of the maubot implementation for Telegram.
func New(token string) (maubot.Bot, error) {
	bot := &TGBot{internal: nil, token: token, uid: maubot.RandomizeUID(), listeners: []chan maubot.Message{}}
	return bot, nil
}

// TGBot is an implementation of maubot for Telegram.
type TGBot struct {
	internal  *telebot.Bot
	listeners []chan maubot.Message
	uid       string
	token     string
	connected bool
	stop      bool
}

// Connect connects to the Telegram servers.
func (bot *TGBot) Connect() error {
	tg, err := telebot.NewBot(bot.token)
	if err != nil {
		return err
	}

	go func() {
		messages := make(chan telebot.Message)
		tg.Listen(messages, 1*time.Second)
		bot.connected = true
		for message := range messages {
			if bot.stop {
				break
			}
			bot.SendToListeners(&TGMessage{bot: bot, internal: message})
			if bot.stop {
				break
			}
		}
		bot.connected = false
	}()
	return nil
}

// UID returns the unique ID for this instance.
func (bot *TGBot) UID() string {
	return bot.uid
}

// Connected returns whether or not the message listener is active.
func (bot *TGBot) Connected() bool {
	return bot.connected
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *TGBot) Disconnect() error {
	bot.stop = true
	return nil
}

// Underlying returns the telebot bot object.
func (bot *TGBot) Underlying() interface{} {
	return bot.internal
}

// SendMessage sends a message to the given channel or user.
func (bot *TGBot) SendMessage(to, message string) {
	bot.internal.SendMessage(nil, message, nil)
}

// SendToListeners ...
func (bot *TGBot) SendToListeners(message maubot.Message) {
	for _, listener := range bot.listeners {
		listener <- message
	}
}

// AddListener adds a message listener
func (bot *TGBot) AddListener(listener chan maubot.Message) {
	bot.listeners = append(bot.listeners, listener)
}
