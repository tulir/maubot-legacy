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

// Package maubot provides interfaces to abstract platforms for chatbots.
package maubot

// Create creates an instance of Maubot
func Create() *Maubot {
	return &Maubot{Interfaces: []Bot{}, Messages: make(chan Message)}
}

// Maubot ...
type Maubot struct {
	Interfaces []Bot
	Messages   chan Message
}

// Add adds an interface to the bot system.
func (mb *Maubot) Add(interfc Bot) {
	mb.Interfaces = append(mb.Interfaces, interfc)
	interfc.AddListener(mb.Messages)
}

// Bot ...
type Bot interface {
	// Underlying returns the underlying API communication system.
	Underlying() interface{}
	// SendMessage sends a message to a room.
	SendMessage(to, message string)
	// AddListener adds a channel the implementation will send messages to.
	AddListener(chan Message)

	// Connect connects to the messaging server.
	Connect() error
	// Connected returns whether or not the bot is connected and ready to use.
	Connected() bool
	// Disconnect disconnects from the messaging server.
	Disconnect() error
}

// Message ...
type Message interface {
	// Underlying returns the underlying API communication system.
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
}
