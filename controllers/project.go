package project

import (
    "html/template"
    "net/http"

    "github.com/gorilla/mux"
    "go-web-app/models"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("../views/*.html"))
}

func SetupRoutes(router *mux.Router) {
    router.HandleFunc("/create-project", CreateProjectHandler).Methods("GET", "POST")
    router.HandleFunc("/projects", ListProjectsHandler).Methods("GET")
}

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "create-project.html", nil)
        return
    }

    name := r.FormValue("name")
    goal := r.FormValue("goal")
    status := r.FormValue("status")
    deadlines := r.FormValue("deadlines")

    newProj := models.Project{
        ProjectName:   name,
        ProjectGoal:   goal,
        ProjectStatus: status,
        Deadlines:     deadlines,
    }

    // TODO: сохранить проект в базу данных
    renderTemplate(w, "create-project.html", newProj)
}

func ListProjectsHandler(w http.ResponseWriter, r *http.Request) {
    projects := make([]models.Project, 0)
    // TODO: получить список проектов из базы данных
    renderTemplate(w, "projects-list.html", projects)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := tpl.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}