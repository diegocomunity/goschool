package commons

import (
	"fmt"
	"net/http"
)

func JS_siteCreator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")

	fmt.Fprint(w, `
	const siteForm = document.getElementById("siteCreatorForm");
	siteForm.addEventListener("submit", (e) => {
		e.preventDefault()
		const data = new FormData(document.getElementById("siteCreatorForm"));
		fetch('/site', {
			body: data,
			method: 'POST'
		}).then(res => res.text())
		.then(res => {
				console.log("RES: ", res)
		})
	})	
	`)
}

func JS_siteCreatorHandler() http.Handler {
	return http.HandlerFunc(JS_siteCreator)
}

func JS_RunProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprint(w, `const runSiteForm = document.getElementById("runSite");
	runSiteForm.addEventListener("submit", (e) => {
		e.preventDefault()
		const data = new FormData(document.getElementById("siteCreatorForm"));
		fetch('/site/run', {
			body: data,
			method: 'POST'
		}).then(res => res.text())
		.then(res => {
				document.getElementById("output").innerHTML = res
				console.log(res)
		})
	})	

	`)
}
func JS_RunProjectHandler() http.Handler {
	return http.HandlerFunc(JS_RunProject)
}
