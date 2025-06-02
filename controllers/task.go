package task

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
    router.HandleFunc("/create-task", CreateTaskHandler).Methods("GET", "POST")
    router.HandleFunc("/tasks/{projectID}", ListTasksHandler).Methods("GET")
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "create-task.html", nil)
        return
    }

    point := r.FormValue("point")
    status := r.FormValue("status")
    deadline := r.FormValue("deadline")
    projectID := r.FormValue("project_id")

    newTask := models.Task{
        TaskPoint:   point,
        TaskStatus:  status,
        Deadline:    deadline,
        ProjectID:   projectID,
    }

    // TODO: сохранить задачу в базу данных
    renderTemplate(w, "create-task.html", newTask)
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    projectID := vars["projectID"]

    tasks := make([]models.Task, 0)
    // TODO: получить список задач для данного проекта из базы данных
    renderTemplate(w, "tasks-list.html", tasks)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := tpl.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}