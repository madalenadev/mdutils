package notification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Telegram interface {
	Info(text string)
	Warning(text string)
	Error(text string)
}

type notificationImpl struct {
	client    *tgbotapi.BotAPI
	context   string
	channelID int64
}

type NotificationTelegramOptions struct {
	ChannelID int64
	Context   string
	Token     string
}

func New(options NotificationTelegramOptions) Telegram {
	if options.Token == "" {
		log.Fatalln("token not provided")
	}

	client, err := tgbotapi.NewBotAPI(options.Token)
	if err != nil {
		log.Fatalln("fail to load tgbotapi")
	}

	pl := &notificationImpl{
		client:    client,
		channelID: options.ChannelID,
		context:   options.Context,
	}

	return pl
}

func (n *notificationImpl) sendText(text string) error {
	msg := tgbotapi.NewMessage(n.channelID, text)
	msg.ParseMode = "html"

	if _, err := n.client.Send(msg); err != nil {
		log.Println("fail to send message")
	}

	return nil
}

func (n *notificationImpl) Info(text string) {
	t := `<strong>🗞 Ino • ` + n.context + `</strong>\n` + text
	n.sendText(t)
}

func (n *notificationImpl) Warning(text string) {
	t := `<strong>⚠️ Warning • ` + n.context + `</strong>\n` + text
	n.sendText(t)
}

func (n *notificationImpl) Error(text string) {
	t := `<strong>🚨 Error • ` + n.context + `</strong>\n` + text
	n.sendText(t)
}
