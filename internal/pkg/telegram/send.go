package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (t *Telegram) Send(html string) error {
	message := tgbotapi.NewMessage(t.channelID, html)
	message.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(message)
	if err != nil {
		return err
	}

	return nil
}
