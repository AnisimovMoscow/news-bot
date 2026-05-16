package app

import "github.com/AnisimovMoscow/news-bot/internal/config"

type App struct {
	config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() {
}
