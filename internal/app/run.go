package app

import (
	"log"
	"slices"
	"strconv"

	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/sports"
)

func (a *App) Run() {
	// получаем все последние
	news, err := sports.LastNews(a.config.TagID, a.config.NewsLimit.All)
	if err != nil {
		log.Println("error", err.Error())
		return
	}

	// сортируем по комментам
	slices.SortFunc(news, func(a, b sports.News) int {
		return b.CommentsCount - a.CommentsCount
	})

	// обрезаем топ
	news = news[:a.config.NewsLimit.Top]

	// сортируем топ по дате
	slices.SortFunc(news, func(a, b sports.News) int {
		return a.PublishedAt.Compare(b.PublishedAt)
	})

	// проверяем новые
	for _, n := range news {
		id, err := strconv.Atoi(n.ID)
		if err != nil {
			log.Println("error", err.Error())
		}

		old, err := a.news.GetByID(id)
		if err != nil {
			log.Println("error", err.Error())
			continue
		}

		if old == nil {
			log.Println("new", n.Title)
			err = a.news.Create(model.News{ID: id})
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

		} else {
			log.Println("old", n.Title)
		}
	}
}
