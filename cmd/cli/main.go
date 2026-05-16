package main

import (
	"log"

	"github.com/AnisimovMoscow/news-bot/internal/app"
	"github.com/AnisimovMoscow/news-bot/internal/config"
	"github.com/AnisimovMoscow/news-bot/internal/database"
)

func main() {
	log.Println("Start")

	cfg := config.New()

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := app.New(cfg, db)
	app.Run()

	log.Println("Done")
}
