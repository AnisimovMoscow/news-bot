package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TagID int
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

	return &Config{
		TagID: tagID,
	}
}
