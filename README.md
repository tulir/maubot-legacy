# Maubot
A chatbot platform abstraction library. If you write a bot that uses Maubot interfaces,
you can easily allow it to connect to any chat platform with Maubot bindings.

Please note that Maubot is at an early stage with only basic support for a few platforms.

Install with `go get maunium.net/go/maubot`

## Usage
First create a `Maubot` object.
```go
var bot = maubot.New()
```

Then create your bots. Most platform bindings want authentication info in `Create()`.
For example, this is how you would connect to Telegram:
```go
import "maunium.net/go/maubot/telegram"
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
* [maubottest](https://github.com/tulir/maubot/tree/master/cmd/maubottest)
* [jesaribot](https://github.com/tulir/jesaribot)

## Supported platforms
#### Currently supported
* [IRC](https://tools.ietf.org/html/rfc1459) with [libmauirc](https://maunium.net/go/libmauirc)
* [Telegram](https://telegram.org/) with [telebot](https://github.com/tucnak/telebot)
* [Slack](https://slack.com) with [slack](https://github.com/nlopes/slack)

#### Support coming soon
* [Matrix](https://matrix.org/)
