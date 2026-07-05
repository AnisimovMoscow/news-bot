package app

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/championat"
	"github.com/AnisimovMoscow/news-bot/internal/pkg/sports"
)

func (a *App) Run() {
	// спортс
	a.sportsNews()

	// чемпионат
	a.championatNews()
}

func (a *App) sportsNews() {
	// получаем все последние
	news, err := sports.LastNews(a.config.Sports.TagID, a.config.NewsLimit.All)
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
	var count int
	for _, n := range news {
		id, err := strconv.Atoi(n.ID)
		if err != nil {
			log.Println("error", err.Error())
		}

		old, err := a.news.GetByID(id, model.SourceSports)
		if err != nil {
			log.Println("error", err.Error())
			continue
		}

		if old == nil {
			// отправляем в канал
			err = a.telegram.Send(getSportsHTML(n))
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			// сохраняем отправленное
			err = a.news.Create(model.News{ID: id}, model.SourceSports)
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			count++
		}
	}

	log.Printf("Sports\ntotal: %d, new:%d\n\n", len(news), count)
}

func print(news []championat.News) {
	for _, n := range news {
		log.Println(n.ID, n.CommentsCount)
	}
}

func (a *App) championatNews() {
	// получаем все последние
	news, err := championat.LastNews(a.config.Championat.Slug)
	if err != nil {
		log.Println("error", err.Error())
		return
	}
	log.Println("all")
	print(news)

	// сортируем по комментам
	slices.SortFunc(news, func(a, b championat.News) int {
		return b.CommentsCount - a.CommentsCount
	})
	log.Println("sort")
	print(news)

	// обрезаем топ
	news = news[:a.config.NewsLimit.Top]
	log.Println("top")
	print(news)

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
			err = a.telegram.Send(getChampionatHTML(n))
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			// сохраняем отправленное
			err = a.news.Create(model.News{ID: id}, model.SourceChampionat)
			if err != nil {
				log.Println("error", err.Error())
				continue
			}

			count++
		}
	}

	log.Printf("Championat\ntotal: %d, new:%d\n\n", len(news), count)
}

func getSportsHTML(news sports.News) string {
	return fmt.Sprintf("%s\n\n<a href=\"%s\">Читать</a>", news.Title, news.URL)
}

func getChampionatHTML(news championat.News) string {
	return fmt.Sprintf("%s\n\n<a href=\"%s\">Читать</a>", news.Title, news.URL)
}
