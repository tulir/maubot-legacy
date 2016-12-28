package slack

import (
	"github.com/nlopes/slack"
	"github.com/satori/go.uuid"
	"maunium.net/go/maubot"
)

func New(token string) (maubot.Bot, error) {
	bot := &SlackBot{token: token, uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}}
	return bot, nil
}

type SlackBot struct {
	internal  *slack.RTM
	listeners []chan maubot.Message
	uid       string
	token     string
}

func (bot *SlackBot) Connect() error {
	client := slack.New(bot.token)
	bot.internal = client.NewRTM()
	go bot.internal.ManageConnection()

	go func() {
		for {
			select {
			case msg := <-bot.internal.IncomingEvents:
				switch ev := msg.Data.(type) {
				case *slack.MessageEvent:
					bot.SendToListeners(&SlackMessage{bot: bot, internal: ev.Msg})
				case *slack.InvalidAuthEvent:
					return
				}
			}
		}
	}()
	return nil
}

// UID returns the unique ID for this instance.
func (bot *SlackBot) UID() string {
	return bot.uid
}

// Connected returns whether or not the message listener is active.
func (bot *SlackBot) Connected() bool {
	return true
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *SlackBot) Disconnect() error {
	err := bot.internal.Disconnect()
	return err
}

// Underlying returns the telebot bot object.
func (bot *SlackBot) Underlying() interface{} {
	return bot.internal
}

// SendMessage sends a message to the given channel or user.
func (bot *SlackBot) SendMessage(to, message string) {
	msg := bot.internal.NewOutgoingMessage(message, to)
	bot.internal.SendMessage(msg)
}

// SendToListeners ...
func (bot *SlackBot) SendToListeners(message maubot.Message) {
	for _, listener := range bot.listeners {
		listener <- message
	}
}

// AddListener adds a message listener
func (bot *SlackBot) AddListener(listener chan maubot.Message) {
	bot.listeners = append(bot.listeners, listener)
}
