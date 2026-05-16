package main

import (
	"log"

	"github.com/AnisimovMoscow/news-bot/internal/app"
	"github.com/AnisimovMoscow/news-bot/internal/config"
)

func main() {
	log.Println("Start")

	cfg := config.New()

	app := app.New(cfg)
	app.Run()

	log.Println("Done")
}
