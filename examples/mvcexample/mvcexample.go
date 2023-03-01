package mvcexample

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegocomunity/goschool/examples/mvcexample/comments"
)

const (
	VERSION = "1"
)

var urlBase_API string
var urlBase_Web string

func init() {
	urlBase_API = "/" //fmt.Sprintf("/api/v%v", VERSION) //result -> api/v1
	urlBase_Web = "/web"
}

type Mvc struct{}

func (Mvc) Run() {

	mux := http.NewServeMux()
	//handler para la api
	mux.Handle(urlBase_API, NewApi())                              // url de la api-> /api/v1
	mux.Handle(urlBase_Web+"/comments", comments.CommentHandler()) //url -> /web/comments
	fmt.Println("run server in port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
