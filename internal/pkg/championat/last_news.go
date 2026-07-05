package championat

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	urlFormat   = "https://www.championat.com/tags/%s/news/"
	host        = "https://www.championat.com"
	commentsApi = "https://comments.rambler.ru/api/app/5/comments-count"
)

var monthReplacer = strings.NewReplacer(
	"января", "1",
	"февраля", "2",
	"марта", "3",
	"апреля", "4",
	"мая", "5",
	"июня", "6",
	"июля", "7",
	"августа", "8",
	"сентября", "9",
	"октября", "10",
	"ноября", "11",
	"декабря", "12",
)

type News struct {
	ID            string
	Xid           string
	Title         string
	CommentsCount int
	PublishedAt   time.Time
	URL           string
}

func LastNews(slug string) ([]News, error) {
	// получаем документ
	url := fmt.Sprintf(urlFormat, slug)
	doc, err := parse(url)
	if err != nil {
		return nil, err
	}

	// собираем результат
	var news []News
	var date string
	doc.Find(".news-items .news-item, .news-items .news-items__head").Each(func(i int, s *goquery.Selection) {
		class, exists := s.Attr("class")
		if !exists {
			return
		}

		if class == "news-items__head" {
			date = monthReplacer.Replace(s.Text())

			return
		}

		xid, exists := s.Find(".news-item__content .news-item__comments .js-comments-count").Attr("data-id")
		if !exists {
			return
		}

		title := s.Find(".news-item__content .news-item__title").Text()

		t := s.Find(".news-item__time").Text()
		published, err := time.Parse("2 1 2006 15:04", fmt.Sprintf("%s %s", date, t))
		if err != nil {
			fmt.Println(err)
			return
		}

		newsUrl, exists := s.Find(".news-item__content .news-item__title").Attr("href")
		if !exists {
			return
		}

		news = append(news, News{
			ID:          strings.TrimLeft(xid, "news_"),
			Xid:         xid,
			Title:       title,
			PublishedAt: published,
			URL:         host + newsUrl,
		})
	})

	// подгружаем комменты
	news, err = loadComments(news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func loadComments(news []News) ([]News, error) {
	params := make([]Param, len(news))
	for i, n := range news {
		params[i] = Param{
			Key:   "xid",
			Value: n.Xid,
		}
	}

	var resp struct {
		Xids map[string]int `json:"xids"`
	}
	err := api(commentsApi, params, &resp)
	if err != nil {
		return nil, err
	}

	for i, n := range news {
		count, ok := resp.Xids[n.Xid]
		if ok {
			news[i].CommentsCount = count
		}
	}

	return news, nil
}
