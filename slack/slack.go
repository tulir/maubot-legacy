package slack

import (
	"github.com/nlopes/slack"
	"github.com/satori/go.uuid"
	"maunium.net/go/maubot-legacy"
)

func New(token string) (maubot.Bot, error) {
	bot := &Bot{token: token, uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}, stop: make(chan bool, 1)}
	return bot, nil
}

type Bot struct {
	internal  *slack.RTM
	listeners []chan maubot.Message
	uid       string
	token     string
	stop      chan bool
}

func (bot *Bot) Connect() error {
	client := slack.New(bot.token)
	bot.internal = client.NewRTM()
	go bot.internal.ManageConnection()

	go func() {
		for {
			select {
			case msg := <-bot.internal.IncomingEvents:
				switch ev := msg.Data.(type) {
				case *slack.MessageEvent:
					bot.SendToListeners(&Message{bot: bot, internal: ev.Msg})
				case *slack.InvalidAuthEvent:
					return
				}
			case stop := <-bot.stop:
				if stop {
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
	return true
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *Bot) Disconnect() error {
	bot.stop <- true
	err := bot.internal.Disconnect()
	return err
}

// Underlying returns the telebot bot object.
func (bot *Bot) Underlying() interface{} {
	return bot.internal
}

// SendMessage sends a message to the given channel or user.
func (bot *Bot) SendMessage(message maubot.OutgoingMessage) {
	bot.internal.SendMessage(
		bot.internal.NewOutgoingMessage(message.Text, message.RoomID))
}

// SendToListeners ...
func (bot *Bot) SendToListeners(message maubot.Message) {
	for _, listener := range bot.listeners {
		listener <- message
	}
}

// AddListener adds a message listener
func (bot *Bot) AddListener(listener chan maubot.Message) {
	bot.listeners = append(bot.listeners, listener)
}
