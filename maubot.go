package maubot

// Create creates an instance of Maubot
func Create() *Maubot {
	return &Maubot{Interfaces: make(map[string]Bot), Messages: make(chan Message)}
}

// Maubot is a wrapper for using multiple messaging platforms.
type Maubot struct {
	Interfaces map[string]Bot
	Messages   chan Message
}

// Add adds an interface to the bot system.
func (mb *Maubot) Add(bot Bot) {
	mb.Interfaces[bot.UID()] = bot
	bot.AddListener(mb.Messages)
}

// Bot is a handler for a single messaging platform.
type Bot interface {
	// Underlying returns the underlying platform bindings.
	// If the handler is implemented within the platform bindings, this can be null.
	Underlying() interface{}
	// SendMessage sends a message to a room.
	SendMessage(to, message string)
	// AddListener adds a channel the implementation will send messages to.
	AddListener(chan Message)
	// UID returns the unique ID for this instance.
	UID() string

	// Connect connects to the messaging server.
	Connect() error
	// Connected returns whether or not the bot is connected and ready to use.
	Connected() bool
	// Disconnect disconnects from the messaging server.
	Disconnect() error
}

// Message is a message (duh)
type Message interface {
	// Underlying returns the underlying platform bindings.
	// If the handler is implemented within the platform bindings, this can be null.
	Underlying() interface{}
	// Text returns the text in the message.
	Text() string
	// Reply sends a message to the room this message originated from.
	Reply(message string)
	// RoomID returns a static room identifier that shouldn't change for the room the message was sent to.
	RoomID() string
	// Room returns the preferred display name for the room the message was sent to.
	Room() string
	// SenderID returns a static user identifier that shouldn't change often.
	SenderID() string
	// Sender returns the preferred display name for the sender.
	Sender() string
	// Source returns the Bot object the message came from.
	Source() Bot
}
