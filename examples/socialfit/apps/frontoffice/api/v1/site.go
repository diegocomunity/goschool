package v1

import (
	"archive/zip"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func Site(w http.ResponseWriter, r *http.Request) {
	var s site
	s.Creator.Title = r.FormValue("title")
	s.Creator.Message = r.FormValue("code")
	s.Name = r.FormValue("name")
	s.MakeSite()
	w.Write([]byte("Creando el sitio"))
	//s.RunSite(w, s.Name)
}
func SiteHandler() http.Handler {
	return http.HandlerFunc(Site)
}

func RunSite(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") + ".zip"
	Run(w, name)

}
func RunSiteHandler() http.Handler {
	return http.HandlerFunc(RunSite)
}

type Creator struct {
	Title, Message string
}
type site struct {
	Name, tpl string
	Creator
}

func (s *site) MakeSite() {
	s.tpl = `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <header class="header">
		<h1>Hello <strong>{{.Message}}</strong></h1>
    </header>
    <main class="main">
		<p>lorem lorem</p>
    </main>
    <footer class="footer">
        <hr>
		<span>FOOTER</span>
    </footer>
</body>
</html>
	`
	t, err := template.New("index.tpl").Parse(s.tpl)
	if err != nil {
		log.Fatalf(err.Error())
	}

	f, err := os.Create(s.Name + ".zip")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//defer f.Close()
	w := zip.NewWriter(f)

	tpl, err := w.Create("index.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = t.Execute(tpl, &s.Creator)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	err = w.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}

}

func Run(w http.ResponseWriter, name string) {
	r, err := zip.OpenReader(name)
	if err != nil {
		w.Write([]byte(`<h1>No hemos encontrado el proyecto!</h1>`))
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
