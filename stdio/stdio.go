// Package stdio provides a maubot implementation for standard input and output
package stdio

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/satori/go.uuid"
	"maunium.net/go/maubot"
)

// New creates an instance of the maubot.Bot implementation for standard IO.
func New() maubot.Bot {
	return &Bot{uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}}
}

// Bot is an implementation of maubot.Bot for standard IO.
type Bot struct {
	listeners []chan maubot.Message
	uid       string
	stop      bool
	connected bool
}

// Connect starts reading stdin
func (bot *Bot) Connect() error {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		bot.connected = true
		for {
			if bot.stop {
				break
			}
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			bot.SendToListeners(&IOMessage{text: strings.TrimSpace(text), bot: bot})
			if bot.stop {
				break
			}
		}
		bot.connected = false
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
	bot.stop = true
	return nil
}

// Underlying returns nil.
func (bot *Bot) Underlying() interface{} {
	return nil
}

// SendMessage sends a message to the given channel or user.
func (bot *Bot) SendMessage(msg maubot.OutgoingMessage) {
	fmt.Println(msg.Text)
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

// IOMessage is an implementation of the Maubot message for standard output.
type IOMessage struct {
	text string
	bot  *Bot
}

// Underlying returns nil.
func (msg *IOMessage) Underlying() interface{} {
	return nil
}

// Source returns the Bot parent of this message.
func (msg *IOMessage) Source() maubot.Bot {
	return msg.bot
}

// Reply prints a message to stdout.
func (msg *IOMessage) Reply(message string) {
	fmt.Print("\n", message, "\n> ")
}

// ReplyWithRef does the same thing as Reply().
func (msg *IOMessage) ReplyWithRef(message string) {
	fmt.Print("\n", message, "\n> ")
}

// Text returns the text in the message.
func (msg *IOMessage) Text() string {
	return msg.text
}

// Room returns "stdio".
func (msg *IOMessage) Room() string {
	return "stdio"
}

// RoomID returns "stdio".
func (msg *IOMessage) RoomID() string {
	return "stdio"
}

// SenderID returns the uid of the user.
func (msg *IOMessage) SenderID() string {
	user, _ := user.Current()
	return user.Uid
}

// Sender returns username of the user.
func (msg *IOMessage) Sender() string {
	user, _ := user.Current()
	return user.Name
}
