package profile

import (
    "html/template"
    "net/http"

    "go-web-app/models"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("../views/*.html"))
}

func SetupRoutes(router *mux.Router) {
    router.HandleFunc("/profile", ProfileHandler).Methods("GET")
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    emp := models.Employee{}
    // TODO: получить текущего пользователя из базы данных
    renderTemplate(w, "profile.html", emp)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := tpl.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}