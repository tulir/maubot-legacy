// Package irc provides a maubot implementation for IRC.
package irc

import (
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
	msg "github.com/sorcix/irc"
	"maunium.net/go/libmauirc"
	"maunium.net/go/maubot-legacy"
)

// New creates an instance of the maubot implementation for Telegram.
func New(nick, user, realname string, addr string) (maubot.Bot, error) {
	addrParts := strings.Split(addr, ":")
	var host string
	var port = 6667
	if len(addrParts) < 2 {
		host = addr
	} else {
		host = addrParts[0]
		port, _ = strconv.Atoi(addrParts[1])
	}
	irc := libmauirc.Create(nick, user, libmauirc.IPv4Address{IP: host, Port: uint16(port)})
	bot := &Bot{internal: irc, nick: nick, user: user, realname: realname, uid: uuid.NewV4().String(), listeners: []chan maubot.Message{}}
	return bot, nil
}

// Bot is an implementation of maubot.Bot for IRC.
type Bot struct {
	internal  libmauirc.Connection
	listeners []chan maubot.Message
	uid       string
	nick      string
	user      string
	realname  string
}

// Connect connects to the IRC server.
func (bot *Bot) Connect() error {
	err := bot.internal.Connect()
	if err != nil {
		return err
	}
	bot.internal.SetRealName(bot.realname)
	bot.internal.AddHandler(msg.PRIVMSG, bot.handlePrivmsg)
	// TODO add more handlers
	return nil
}

// UID returns the unique ID for this instance.
func (bot *Bot) UID() string {
	return bot.uid
}

// Connected returns whether or not the message listener is active.
func (bot *Bot) Connected() bool {
	return bot.internal.Connected()
}

// Disconnect stops listening for messages. It may or may not actually disconnect.
func (bot *Bot) Disconnect() error {
	bot.internal.Quit()
	go func() {
		if bot.Connected() {
			bot.internal.Disconnect()
		}
	}()
	return nil
}

// Underlying returns the telebot bot object.
func (bot *Bot) Underlying() interface{} {
	return bot.internal
}

// SendMessage sends a message to the given channel or user.
func (bot *Bot) SendMessage(msg maubot.OutgoingMessage) {
	bot.internal.Privmsg(msg.RoomID, msg.Text)
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
