package page_go

import (
    "net/http"
	"html/template"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}
