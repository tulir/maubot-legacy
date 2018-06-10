// Package matrix provides a maubot implementation for Matrix.
package matrix

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"maunium.net/go/maubot-legacy"
	"maunium.net/go/mautrix"
)

// New creates an instance of the maubot.Bot implementation for Matrix.
func New(homeserver, user, password string) (maubot.Bot, error) {
	bot := &Bot{internal: mautrix.Create(homeserver), user: user, password: password,
		uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}, stop: make(chan bool, 1)}
	return bot, nil
}

// Bot is an implementation of maubot.Bot for Telegram.
type Bot struct {
	internal  *mautrix.Session
	listeners []chan maubot.Message
	uid       string
	token     string
	connected bool
	user      string
	password  string
	stop      chan bool
}

// Connect connects to the Matrix homeserver.
func (bot *Bot) Connect() error {
	err := bot.internal.PasswordLogin(bot.user, bot.password)
	if err != nil {
		return err
	}

	go func() {
		bot.connected = true
		go bot.internal.Listen()

		for {
			select {
			case evt := <-bot.internal.Timeline:
				if evt.Sender == bot.internal.MatrixID {
					continue
				}
				evt.Age = (time.Now().UnixNano() / 1000000) - evt.OriginServerTime
				const tenSeconds = 10 * 1000
				if evt.Age > tenSeconds {
					continue
				}
				switch evt.Type {
				case mautrix.EvtRoomMessage:
					bot.SendToListeners(&Message{internal: evt, bot: bot})
				default:
					fmt.Println(evt.Type)
				}
			case roomID := <-bot.internal.InviteChan:
				invite := bot.internal.Invites[roomID]
				fmt.Printf("%s invited me to %s (%s)\n", invite.Sender, invite.Name, invite.ID)
				fmt.Println(invite.Accept())
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
	bot.sendMessage(msg.RoomID, msg.Text)
}

func (bot *Bot) sendMessage(roomID, text string) {
	room := bot.internal.Rooms[roomID]
	if room != nil {
		room.Send(text)
	}
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
