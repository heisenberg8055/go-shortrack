package template

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Name string
}

func Template(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").ParseGlob("./internal/templates/resources/*"))
	tmpl.ExecuteTemplate(w, "index.html", PageData{
		Name: "Test",
	})
}
