package app

import (
	"log"
	"slices"
	"strconv"

	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/championat"
)

func (a *App) championatNews() {
	// получаем все последние
	news, err := championat.LastNews(a.config.Championat.Slug)
	if err != nil {
		log.Println("error", err.Error())
		return
	}

	// сортируем по комментам
	slices.SortFunc(news, func(a, b championat.News) int {
		return b.CommentsCount - a.CommentsCount
	})

	// обрезаем топ
	news = news[:a.config.NewsLimit.Top]

	// сортируем топ по дате
	slices.SortFunc(news, func(a, b championat.News) int {
		return a.PublishedAt.Compare(b.PublishedAt)
	})

	// проверяем новые
	var count int
	for _, n := range news {
		id, err := strconv.Atoi(n.ID)
		if err != nil {
			log.Println("error", err.Error())
		}

		old, err := a.news.GetByID(id, model.SourceChampionat)
		if err != nil {
			log.Println("error", err.Error())
			continue
		}

		if old == nil {
			// отправляем в канал
			err = a.telegram.Send(getHTML(n.Title, n.URL))
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			// сохраняем отправленное
			err = a.news.CreateID(model.News{ID: id}, model.SourceChampionat)
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			count++
		}
	}

	log.Printf("Championat\ntotal: %d, new:%d\n\n", len(news), count)
}
