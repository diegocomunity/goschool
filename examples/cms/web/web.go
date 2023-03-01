package web

import (
	"archive/zip"
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

//go:embed public
var web embed.FS

func Web(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("Content-Type") == "text/json; charset=utf-8" {
		http.Error(w, "No content", http.StatusNoContent)
		return
	}

	/*
		f, err := web.Open("public/foo.html")
		if err != nil {
			log.Fatalf(err.Error())
		}
		_, err = io.Copy(w, f)
		if err != nil {
			log.Fatalf(err.Error())
		}
	*/
	t, err := template.New("index.tpl").ParseFS(web, "public/index.tpl", "public/foo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, err.Error(), http.StatusNoContent)
	}
}
func WebHandler() http.Handler {
	return http.HandlerFunc(Web)
}

type menuModel struct {
	Items []string
}

//type webModel struct {
//}

func BuildProjet(w http.ResponseWriter, r *http.Request) {
	option := r.FormValue("option")
	id := r.FormValue("id")
	if option == "execute" {
		execute(w, id+".zip")
	}
	if id != "" {
		jsonStream := r.FormValue("menu")
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		for {
			var menu menuModel
			if err := dec.Decode(&menu); err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf(err.Error())
			}
			println(menu.Items[0])
			go build(r.FormValue("id"), menu)
		}
	}
	io.WriteString(w, "Id es un campo requerido")
}
func BuildProjetHandler() http.Handler {
	return http.HandlerFunc(BuildProjet)
}
func build(id string, menu menuModel) {
	var test = struct {
		Title string
		Items []string
	}{
		Title: "Esto es un ejemplo de algùn titulo",
		Items: menu.Items,
	}
	t, err := template.New("template.tpl").ParseFS(web, "public/template.tpl")
	if err != nil {
		log.Fatalf(err.Error())
	}

	f, err := os.Create(id + ".zip")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//defer f.Close()
	w := zip.NewWriter(f)

	tpl, err := w.Create("index.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(tpl, test)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	err = w.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}

}

func execute(w http.ResponseWriter, id string) {
	r, err := zip.OpenReader(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatalf(err.Error())
		}
		_, err = io.Copy(w, rc)
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer rc.Close()
	}
}
