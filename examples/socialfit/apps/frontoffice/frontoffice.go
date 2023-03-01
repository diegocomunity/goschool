package frontoffice

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/diegocomunity/goschool/examples/socialfit/apps/frontoffice/api/v1"
	"github.com/diegocomunity/goschool/examples/socialfit/commons"
)

const (
	VERSION = "v1"
	PORT    = ":8080"
)

var api http.Handler

func init() {
	if VERSION != "v1" {
		panic("por el momento estamos en la version 1")
	}
	api = v1.New() //instancia la api en la versiòn 1
}
func Bootstrap() {
	mux := http.NewServeMux()
	mux.Handle("/", api)
	mux.Handle("/css/global", commons.CSS_GlobalHandler())
	mux.Handle("/js/sitecreator", commons.JS_siteCreatorHandler())
	mux.Handle("/js/run", commons.JS_RunProjectHandler())
	fmt.Println("run server in port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
