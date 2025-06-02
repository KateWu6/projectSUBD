package auth

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
    router.HandleFunc("/register", RegisterHandler).Methods("GET", "POST")
    router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "register.html", nil)
        return
    }

    // Обработка формы регистрации
    username := r.FormValue("username")
    email := r.FormValue("email")
    password := r.FormValue("password")

    // Примерная проверка и сохранение в БД
    newUser := models.User{
        Username: username,
        Email:    email,
        Password: password,
    }

    // TODO: сохранить пользователя в базу данных
    renderTemplate(w, "register.html", newUser)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "login.html", nil)
        return
    }

    // Обработка формы авторизации
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Примерная проверка и авторизация пользователя
    // TODO: проверка в базе данных
    renderTemplate(w, "login.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := tpl.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}