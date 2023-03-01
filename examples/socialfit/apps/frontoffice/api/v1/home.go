package v1

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed web
var home embed.FS

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tpl").ParseFS(home, "web/index.tpl")
	if err != nil {
		panic("error wiht template")
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalf(err.Error())
	}
}
func HomeHandler() http.Handler {
	return http.HandlerFunc(Home)
}
