package news

import (
	"fmt"

	"github.com/AnisimovMoscow/news-bot/internal/model"
)

func (r *Repository) CreateURL(news model.News, source model.Source) error {
	table := tableName[source]
	query := fmt.Sprintf("INSERT INTO %s (url) VALUES (?)", table)
	_, err := r.db.Exec(query, news.URL)
	if err != nil {
		return err
	}

	return nil
}
