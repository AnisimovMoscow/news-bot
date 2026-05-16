package telegram

import (
	"log"

	"github.com/AnisimovMoscow/news-bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot       *tgbotapi.BotAPI
	channelID int64
}

func New(cfg config.Telegram) *Telegram {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	return &Telegram{
		bot:       bot,
		channelID: cfg.Channel,
	}
}
