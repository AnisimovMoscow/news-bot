package telegram

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/AnisimovMoscow/news-bot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"
)

const timeout = 30 * time.Second

type Telegram struct {
	bot       *tgbotapi.BotAPI
	channelID int64
}

func New(cfg config.Telegram) *Telegram {
	client := &http.Client{
		Timeout: timeout,
	}

	if cfg.Proxy != "" {
		u, err := url.Parse(cfg.Proxy)
		if err != nil {
			log.Fatal(err)
		}

		dialer, err := proxy.FromURL(u, proxy.Direct)
		if err != nil {
			log.Fatal(err)
		}

		client.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}
	}

	bot, err := tgbotapi.NewBotAPIWithClient(cfg.Token, tgbotapi.APIEndpoint, client)
	if err != nil {
		log.Fatal(err)
	}

	return &Telegram{
		bot:       bot,
		channelID: cfg.Channel,
	}
}
