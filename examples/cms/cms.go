package cms

import (
	"log"
	"net/http"
)

func New() *cms {
	return &cms{}
}

type cms struct{}

func (c *cms) Run() {
	startRoutes()
}

func startRoutes() {
	mux := http.NewServeMux()
	mux.Handle("/", api{})
	println("run serve in port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
