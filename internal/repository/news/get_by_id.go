package news

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AnisimovMoscow/news-bot/internal/model"
)

func (r *Repository) GetByID(id int, source model.Source) (*model.News, error) {
	var news model.News
	table := tableName[source]
	query := fmt.Sprintf("SELECT id FROM %s WHERE id = ?", table)
	err := r.db.Get(&news, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &news, nil
}
