// Package stdio provides a maubot implementation for standard input and output
package stdio

import (
	"bufio"
	"fmt"
	"maunium.net/go/maubot"
	"os"
	"os/user"
	"strings"
)

// New creates an instance of the maubot implementation for standard input and output.
func New() maubot.Bot {
	return &IOBot{uid: maubot.RandomizeUID(), listeners: []chan maubot.Message{}}
}

// IOBot is an implementation of maubot for standard input and output.
type IOBot struct {
	listeners []chan maubot.Message
	uid       string
	stop      bool
	connected bool
}

// Connect starts reading stdin
func (bot *IOBot) Connect() error {
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
func (bot *IOBot) UID() string {
	return bot.uid
}

// Connected returns whether or not the message listener is active.
func (bot *IOBot) Connected() bool {
	return bot.connected
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *IOBot) Disconnect() error {
	bot.stop = true
	return nil
}

// Underlying returns nil.
func (bot *IOBot) Underlying() interface{} {
	return nil
}

// SendMessage sends a message to the given channel or user.
func (bot *IOBot) SendMessage(to, message string) {
	fmt.Println(message)
}

// SendToListeners ...
func (bot *IOBot) SendToListeners(message maubot.Message) {
	for _, listener := range bot.listeners {
		listener <- message
	}
}

// AddListener adds a message listener
func (bot *IOBot) AddListener(listener chan maubot.Message) {
	bot.listeners = append(bot.listeners, listener)
}

// IOMessage is an implementation of the Maubot message for standard output.
type IOMessage struct {
	text string
	bot  *IOBot
}

// Underlying returns nil.
func (msg *IOMessage) Underlying() interface{} {
	return nil
}

// Source returns the IOBot parent of this message.
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
