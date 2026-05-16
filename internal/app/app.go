package app

import (
	"log"
	"slices"

	"github.com/AnisimovMoscow/news-bot/internal/config"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/sports"
)

const (
	allNewsLimit = 50
	topNewsLimit = 15
)

type App struct {
	config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Run() {
	// получаем все последние
	news, err := sports.LastNews(a.config.TagID, allNewsLimit)
	if err != nil {
		log.Println(err)
		return
	}

	// сортируем
	slices.SortFunc(news, func(a, b sports.News) int {
		return b.CommentsCount - a.CommentsCount
	})

	// обрезаем топ
	news = news[:topNewsLimit]

	// сортируем топ
	slices.SortFunc(news, func(a, b sports.News) int {
		return b.PublishedAt.Compare(a.PublishedAt)
	})

	for _, n := range news {
		log.Println(n)
	}
}
