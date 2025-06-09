package main

import (
    "net/http"
    "projectSUBD/page_go"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // Маршруты
    router.HandleFunc("/", page_go.HomeHandler)

     router.HandleFunc("/login", page_go.LoginHandler)

    

    router.HandleFunc("/projects/", page_go.ProjectsHandler)
    router.HandleFunc("/users/", page_go.UsersHandler)

    router.HandleFunc("/tasks/{username}", page_go.TasksHandler)
    router.HandleFunc("/projects/{projectID}/tasks", page_go.ProjectTasksHandler)

    router.HandleFunc("/add-task", page_go.AddTaskHandler).Methods("GET", "POST")

    // Запускаем сервер
    http.ListenAndServe(":8080", router)
}
