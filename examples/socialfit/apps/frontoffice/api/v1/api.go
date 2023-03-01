package v1

import (
	"io"
	"net/http"
)

func New() *api {
	return &api{}
}

type api struct{}

func (api *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch url := r.URL.RequestURI(); url {
	case "/":
		Home(w, r)
		return
	case "/site":
		Site(w, r)
		return
	case "/site/run":
		RunSite(w, r)
		return
	}
	io.WriteString(w, "fobar")
}
