package bobsoccer

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	urlFormat = "https://bobsoccer.ru/tags/%s.html"
	host      = "https://bobsoccer.ru"
)

type News struct {
	Title         string
	CommentsCount int
	PublishedAt   time.Time
	URL           string
	RelativeURL   string
}

func LastNews(slug string, limit int) ([]News, error) {
	// получаем документ
	url := fmt.Sprintf(urlFormat, slug)

	// собираем результат
	var news []News
	page := 1
	for len(news) < limit && page <= 10 {
		doc, err := parse(url, []Param{{Key: "part", Value: strconv.Itoa(page)}})
		if err != nil {
			return nil, err
		}

		doc.Find(".middle-side .blog-list").Each(func(_ int, s *goquery.Selection) {
			title := s.Find("h2 a").Text()

			var count int
			grabNext := false
			s.Find(".element-stat").Contents().Each(func(_ int, s *goquery.Selection) {
				if grabNext {
					text := strings.TrimSpace(s.Text())
					if text != "" {
						count, err = strconv.Atoi(text)
						if err != nil {
							fmt.Println(err)
							return
						}
						grabNext = false
					}
				}
				if s.HasClass("fa-comment-o") {
					grabNext = true
				}
			})

			datetime, exists := s.Find(".element-options time").Attr("datetime")
			if !exists {
				return
			}

			var published time.Time
			sec, err := strconv.ParseInt(datetime, 10, 64)
			if err == nil {
				published = time.Unix(sec, 0).UTC()
			} else {
				published, err = time.Parse("2006-01-02 15:04:05", datetime)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			newsUrl, exists := s.Find("h2 a").Attr("href")
			if !exists {
				return
			}

			news = append(news, News{
				Title:         title,
				CommentsCount: count,
				PublishedAt:   published,
				URL:           host + newsUrl,
				RelativeURL:   newsUrl,
			})
		})

		page++
	}

	return news, nil
}
