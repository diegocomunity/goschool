package website

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed testdata
var content embed.FS

type C struct {
	Nick, Message string
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.New("index.html").ParseFS(content, "testdata/index.html")
	if err != nil {
		log.Fatal(err)
	}
	var message = new(C)
	message.Nick = "Nick por defecto"
	message.Message = "COMENTARIO"
	if r.Method == "POST" {
		println("reciviendo mensaje")
		message.Nick = r.FormValue("nick")
		message.Message = r.FormValue("message")
		t, err := template.New("index.html").ParseFS(content, "testdata/index.html")
		if err != nil {
			log.Fatalf(err.Error())
		}
		if err := t.Execute(w, message); err != nil {
			log.Fatalf("error mi hermano")
		}
		println("terminando mensaje")
		return
	}
	if err := t.Execute(w, message); err != nil {
		log.Fatal(err)
	}
}

func HomeHandler() http.Handler {
	return http.HandlerFunc(Home)
}
