package app

import (
	"log"
	"slices"

	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/sport24"
)

func (a *App) sport24News() {
	// получаем все последние
	news, err := sport24.LastNews(a.config.Sport24.TagID, a.config.NewsLimit.All)
	if err != nil {
		log.Println("error", err.Error())
		return
	}

	// сортируем по комментам
	slices.SortFunc(news, func(a, b sport24.News) int {
		return b.CommentsCount - a.CommentsCount
	})

	// обрезаем топ
	news = news[:a.config.NewsLimit.Top]

	// сортируем топ по дате
	slices.SortFunc(news, func(a, b sport24.News) int {
		return a.PublishedAt.Time.Compare(b.PublishedAt.Time)
	})

	// проверяем новые
	var count int
	for _, n := range news {
		if n.CommentsCount == 0 {
			continue
		}
		old, err := a.news.GetByID(n.ID, model.SourceSport24)
		if err != nil {
			log.Println("error", err.Error())
			continue
		}

		if old == nil {
			// отправляем в канал
			err = a.telegram.Send(getHTML(n.Title, n.URL()))
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			// сохраняем отправленное
			err = a.news.Create(model.News{ID: n.ID}, model.SourceSport24)
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			count++
		}
	}

	log.Printf("Sport24\ntotal: %d, new:%d\n\n", len(news), count)
}
