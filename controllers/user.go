package user

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
    router.HandleFunc("/create-user", CreateUserHandler).Methods("GET", "POST")
    router.HandleFunc("/users", ListUsersHandler).Methods("GET")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        renderTemplate(w, "create-user.html", nil)
        return
    }

    surname := r.FormValue("surname")
    firstname := r.FormValue("firstname")
    middlename := r.FormValue("middlename")
    passportSeries := r.FormValue("passport_series")
    passportNumber := r.FormValue("passport_number")
    post := r.FormValue("post")
    contacts := r.FormValue("contacts")
    photo := r.FormValue("photo")
    rolID := r.FormValue("rol_id")
    addressID := r.FormValue("address_id")
    divID := r.FormValue("division_id")

    newEmp := models.Employee{
        Surname:                surname,
        Firstname:              firstname,
        Middlename:             middlename,
        PassportSeries:         passportSeries,
        PassportNumber:         passportNumber,
        Post:                   post,
        Contacts:               contacts,
        Photo:                  photo,
        RolID:                  rolID,
        AddressID:              addressID,
        StructuralDivisionsID:  divID,
    }

    // TODO: сохранить сотрудника в базу данных
    renderTemplate(w, "create-user.html", newEmp)
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    employees := make([]models.Employee, 0)
    // TODO: получить список сотрудников из базы данных
    renderTemplate(w, "users-list.html", employees)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := tpl.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}