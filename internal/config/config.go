package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TagID     int
	NewsLimit NewsLimit
	Telegram  Telegram
	DB        string
}
type NewsLimit struct {
	All int
	Top int
}
type Telegram struct {
	Token   string
	Channel int64
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	tagID, err := strconv.Atoi(os.Getenv("TAG_ID"))
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
		TagID: tagID,
		NewsLimit: NewsLimit{
			All: allNewsLimit,
			Top: topNewsLimit,
		},
		Telegram: Telegram{
			Token:   os.Getenv("TELEGRAM_TOKEN"),
			Channel: channelID,
		},
		DB: os.Getenv("DB"),
	}
}
