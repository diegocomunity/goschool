package mvcexample

import (
	"net/http"

	"github.com/diegocomunity/goschool/examples/mvcexample/comments"
	"github.com/diegocomunity/goschool/examples/mvcexample/middleware"
	"github.com/diegocomunity/goschool/examples/mvcexample/website"
)

func NewApi() *api {
	return &api{}
}

type api struct{}

func (api *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path para obtener cualquier peticiòn por cualquier metodo
	switch {
	case r.URL.RequestURI() == "/":
		website.Home(w, r)
		return
	case r.URL.RequestURI() == "/api/comments":
		//para retornar un JSON en el handler agregar en la cabecera Content-Type text/json o text/plain
		w.Header().Set("Content-Type", "text/json; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		comments.Comments(w, r)
		return
	default:
		middleware.NotFound(w, r)
		return
	}
	//io.WriteString(w, "From the API /")
}
