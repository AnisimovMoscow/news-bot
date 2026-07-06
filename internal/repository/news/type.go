package news

import (
	"github.com/AnisimovMoscow/news-bot/internal/model"
	"github.com/jmoiron/sqlx"
)

var tableName = map[model.Source]string{
	model.SourceSports:     "sports_news",
	model.SourceChampionat: "championat_news",
	model.SourceSport24:    "sport24_news",
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
