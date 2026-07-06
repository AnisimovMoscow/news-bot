package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Sports     Sports
	Championat Championat
	Sport24    Sport24
	NewsLimit  NewsLimit
	Telegram   Telegram
	DB         string
}

type Sports struct {
	TagID int
}

type Championat struct {
	Slug string
}

type Sport24 struct {
	TagID int
}

type NewsLimit struct {
	All int
	Top int
}

type Telegram struct {
	Token   string
	Channel int64
	Proxy   string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	sportsTagID, err := strconv.Atoi(os.Getenv("SPORTS_TAG_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}

	sport24TagID, err := strconv.Atoi(os.Getenv("SPORT24_TAG_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}

	allNewsLimit, err := strconv.Atoi(os.Getenv("NEWS_LIMIT_ALL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	topNewsLimit, err := strconv.Atoi(os.Getenv("NEWS_LIMIT_TOP"))
	if err != nil {
		log.Fatal(err.Error())
	}

	channelID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHANNEL_ID"), 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Config{
		Sports: Sports{
			TagID: sportsTagID,
		},
		Championat: Championat{
			Slug: os.Getenv("CHAMPIONAT_SLUG"),
		},
		Sport24: Sport24{
			TagID: sport24TagID,
		},
		NewsLimit: NewsLimit{
			All: allNewsLimit,
			Top: topNewsLimit,
		},
		Telegram: Telegram{
			Token:   os.Getenv("TELEGRAM_TOKEN"),
			Channel: channelID,
			Proxy:   os.Getenv("TELEGRAM_PROXY"),
		},
		DB: os.Getenv("DB"),
	}
}
