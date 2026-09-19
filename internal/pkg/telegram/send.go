package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) Send(html string) error {
	if t == nil {
		log.Println("Telegram Send", html)
		return nil
	}

	message := tgbotapi.NewMessage(t.channelID, html)
	message.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(message)
	if err != nil {
		return err
	}

	return nil
}
