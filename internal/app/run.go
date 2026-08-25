package app

import (
	"fmt"
)

func (a *App) Run() {
	// спортс
	//a.sportsNews()

	// чемпионат
	//a.championatNews()

	// sport24
	//a.sport24News()

	// bobsoccer
	a.bobsoccerNews()
}

func getHTML(title, url string) string {
	return fmt.Sprintf("%s\n\n<a href=\"%s\">Читать</a>", title, url)
}
