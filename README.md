# No longer maintained
[Matrix](https://matrix.org) has bridges, so I'm focusing on making bridge-friendly
Matrix bots rather than trying to make bots that speak many protocols. I'm also bad
at naming, so I renamed this project to maubot-legacy in order to use [maubot](https://github.com/tulir/maubot)
for my Matrix bot system.

Matrix bridges that can replace this library:
* [mautrix-telegram](https://github.com/tulir/mautrix-telegram)
* [matrix-appservice-irc](https://github.com/matrix-org/matrix-appservice-irc)
* [matrix-appservice-slack](https://github.com/matrix-org/matrix-appservice-slack)

# Maubot (Legacy)
A chatbot platform abstraction library. If you write a bot that uses Maubot interfaces,
you can easily allow it to connect to any chat platform with Maubot bindings.

Please note that Maubot is at an early stage with only basic support for a few platforms.

Install with `go get maunium.net/go/maubot-legacy`

## Usage
First create a `Maubot` object.
```go
var bot = maubot.New()
```

Then create your bots. Most platform bindings want authentication info in `Create()`.
For example, this is how you would connect to Telegram:
```go
import "maunium.net/go/maubot-legacy/telegram"
...
// Initialize the platform binding object.
var tgBot = telegram.New("botToken")
// Add the initialized platform binding object to the Maubot wrapper object.
bot.Add(tgBot)
// Actually connect to the chat platform.
//
// This doesn't necessarily have to be after bot.Add(), but there's a chance
// that you'll miss some messages if you connect before adding the object to the
// Maubot wrapper.
tgBot.Connect()
```

Finally listen to incoming messages and handle them.
```go
import "fmt"
...
for message := range bot.Messages() {
  fmt.Println("Received \"%s\" from %s", message.Text(), message.Sender())
}
```

## Actual examples
* [maubottest](https://github.com/tulir/maubot-legacy/tree/master/cmd/maubottest)
* [jesaribot](https://github.com/tulir/jesaribot)

## Supported platforms
* [IRC](https://tools.ietf.org/html/rfc1459) with [libmauirc](https://maunium.net/go/libmauirc)
* [Telegram](https://telegram.org/) with [telebot](https://github.com/tucnak/telebot)
* [Slack](https://slack.com) with [slack](https://github.com/nlopes/slack)
* [Matrix](https://matrix.org) with [mautrix](https://maunium.net/go/mautrix)
