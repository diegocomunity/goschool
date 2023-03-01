package middleware

import (
	"html/template"
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<h1>No hemos encontrado este sitio</h1>
	`
	t, err := template.New("404").Parse(tpl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = t.Execute(w, t)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(NotFound)
}
