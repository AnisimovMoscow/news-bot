package app

import (
	"log"
	"slices"

	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/bobsoccer"
)

func (a *App) bobsoccerNews() {
	// получаем все последние
	news, err := bobsoccer.LastNews(a.config.Bobsoccer.Slug, 50)
	if err != nil {
		log.Println("error", err.Error())
		return
	}

	// сортируем по комментам
	slices.SortFunc(news, func(a, b bobsoccer.News) int {
		return b.CommentsCount - a.CommentsCount
	})

	// обрезаем топ
	news = news[:a.config.NewsLimit.Top]

	// сортируем топ по дате
	slices.SortFunc(news, func(a, b bobsoccer.News) int {
		return a.PublishedAt.Compare(b.PublishedAt)
	})

	// проверяем новые
	var count int
	for _, n := range news {
		old, err := a.news.GetByURL(n.RelativeURL, model.SourceBobsoccer)
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
			err = a.news.CreateURL(model.News{URL: n.RelativeURL}, model.SourceBobsoccer)
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			count++
		}
	}

	log.Printf("Bobsoccer\ntotal: %d, new:%d\n\n", len(news), count)
}
