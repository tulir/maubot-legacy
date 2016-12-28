package maubot

// Maubot is a wrapper for simultaneously using multiple chat interfaces that implement the Bot interface.
type Maubot interface {
	// Add adds a chat interface and registers this Maubot wrapper as a listener.
	Add(bot Bot)
	// Get returns the chat interface with the given UID.
	Get(uid string) Bot
	// Remove removes a chat interface from this Maubot wrapper.
	Remove(uid string)
	// Messages returns the message channel for this Maubot wrapper.
	Messages() chan Message
	// SendMessage sends an OutgoingMessage to the correct messaging platform and room.
	SendMessage(msg OutgoingMessage)
}

// Bot is a handler for a single messaging platform.
type Bot interface {
	// Underlying returns the underlying platform bindings.
	// If the handler is implemented within the platform bindings, this can be null.
	Underlying() interface{}
	// SendMessage sends an OutgoingMessage to the correct room.
	SendMessage(msg OutgoingMessage)
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

// Message is an incoming message.
type Message interface {
	// Underlying returns the underlying platform bindings.
	// If the handler is implemented within the platform bindings, this can be null.
	Underlying() interface{}
	// Text returns the text in the message.
	Text() string
	// Reply sends a message to the room this message originated from.
	Reply(message string)
	// ReplyWithRef sends a message to the room this message originated from with
	// a reference to the original sender.
	ReplyWithRef(message string)
	// RoomID returns a static room identifier that shouldn't change for the room
	// the message was sent to.
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

// OutgoingMessage is a message that's going to be sent by the bot.
type OutgoingMessage struct {
	// Text is the text in the message.
	Text string
	// RoomID is the static identifier for the room the message should be sent to.
	RoomID string
	// PlatformID is the UID of the chat interface the message should be sent to.
	PlatformID string
}
