package app

import (
	"fmt"
)

func (a *App) Run() {
	// спортс
	a.sportsNews()

	// чемпионат
	a.championatNews()

	// спорт24
	a.sport24News()
}

func getHTML(title, url string) string {
	return fmt.Sprintf("%s\n\n<a href=\"%s\">Читать</a>", title, url)
}
