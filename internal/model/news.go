package model

type News struct {
	ID  int    `db:"id"`
	URL string `db:"url"`
}
