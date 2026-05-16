package news

import "github.com/AnisimovMoscow/news-bot/internal/model"

func (r *Repository) Create(news model.News) error {
	_, err := r.db.Exec("INSERT INTO news (id) VALUES (?)", news.ID)
	if err != nil {
		return err
	}

	return nil
}
