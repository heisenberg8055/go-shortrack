package template

import (
	"html/template"
	"net/http"
)

func RenderHome(w http.ResponseWriter) {
	tp, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tp.Execute(w, nil)
}

func RenderHomeShortLink(w http.ResponseWriter, shorturl string) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, shorturl)
}
