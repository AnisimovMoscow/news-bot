package app

import (
	"github.com/AnisimovMoscow/news-bot/internal/config"
	"github.com/AnisimovMoscow/news-bot/internal/repository/news"
	"github.com/jmoiron/sqlx"
)

type App struct {
	config *config.Config
	news   *news.Repository
}

func New(cfg *config.Config, db *sqlx.DB) *App {
	repo := news.New(db)

	return &App{
		config: cfg,
		news:   repo,
	}
}
