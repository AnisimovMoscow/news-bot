package news

import (
	"fmt"

	"github.com/AnisimovMoscow/news-bot/internal/model"
)

func (r *Repository) Create(news model.News, source model.Source) error {
	table := tableName[source]
	query := fmt.Sprintf("INSERT INTO %s (id) VALUES (?)", table)
	_, err := r.db.Exec(query, news.ID)
	if err != nil {
		return err
	}

	return nil
}
