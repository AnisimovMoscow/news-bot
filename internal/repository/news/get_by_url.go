package news

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AnisimovMoscow/news-bot/internal/model"
)

func (r *Repository) GetByURL(url string, source model.Source) (*model.News, error) {
	var news model.News
	table := tableName[source]
	query := fmt.Sprintf("SELECT url FROM %s WHERE url = ?", table)
	err := r.db.Get(&news, query, url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &news, nil
}
