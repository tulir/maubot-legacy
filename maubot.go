package maubot

type maubotImpl struct {
	interfaces map[string]Bot
	messages   chan Message
}

// New creates an object that implements the Maubot interface.
func New() Maubot {
	return &maubotImpl{interfaces: make(map[string]Bot), messages: make(chan Message)}
}

func (mb *maubotImpl) Add(bot Bot) {
	mb.interfaces[bot.UID()] = bot
	bot.AddListener(mb.messages)
}

func (mb *maubotImpl) Get(uid string) Bot {
	return mb.interfaces[uid]
}

func (mb *maubotImpl) Remove(uid string) {
	delete(mb.interfaces, uid)
}

func (mb *maubotImpl) Messages() chan Message {
	return mb.messages
}

func (mb *maubotImpl) SendMessage(msg OutgoingMessage) {

}
