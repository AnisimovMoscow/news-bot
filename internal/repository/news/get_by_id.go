package news

import (
	"database/sql"
	"errors"

	"github.com/AnisimovMoscow/news-bot/internal/model"
)

func (r *Repository) GetByID(id int) (*model.News, error) {
	var news model.News
	err := r.db.Get(&news, "SELECT id FROM news WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &news, nil
}
