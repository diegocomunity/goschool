package cms

import (
	"io"
	"net/http"
	"os"

	"github.com/diegocomunity/goschool/examples/cms/web"
)

type api struct{}

func (api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		switch r.URL.RequestURI() {
		case "/api/projet/build":
			web.BuildProjet(w, r)
			return
		}
	}
	if r.Method == "GET" {
		switch r.URL.RequestURI() {
		case "/":
			w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
			w.WriteHeader(http.StatusOK)
			web.Web(w, r)
			return
		case "/projets":
			file, err := os.ReadFile("examples/cms/web/public/foo.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(file)
			return
		default:
			http.NotFound(w, r)
			return
		}
	}
	io.WriteString(w, "run serve")
}
