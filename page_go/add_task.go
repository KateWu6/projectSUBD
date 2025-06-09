package page_go

import (
    "net/http"
    "log"
    "fmt"
	"html/template"
    "projectSUBD/bd" // Assuming bd package is in the same module
)

// Функция обработки страниц добавления задачи
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Если приходит GET-запрос, показываем форму для добавления задачи
        showAddTaskForm(w, r)
    case http.MethodPost:
        // Если пришёл POST-запрос, добавляем задачу
        handleAddTask(w, r)
    default:
        // Если другие методы, возвращаем ошибку
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Метод %s не поддерживается.", r.Method)
    }
}

// Отображение пустой формы для добавления задачи
func showAddTaskForm(w http.ResponseWriter, r *http.Request) {
    db, err := bd.Connect()
    if err != nil {
        log.Println("Ошибка подключения к БД:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    employeesRows, err := db.Query("SELECT username FROM user_info_view")
    if err != nil {
        log.Println("Ошибка загрузки сотрудников:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer employeesRows.Close()

    projectsRows, err := db.Query("SELECT project_name FROM projects")
    if err != nil {
        log.Println("Ошибка загрузки проектов:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer projectsRows.Close()

    var employees []string
    var projects []string

    for employeesRows.Next() {
        var emp string
        err := employeesRows.Scan(&emp)
        if err != nil {
            log.Println("Ошибка чтения сотрудника:", err)
            continue
        }
        employees = append(employees, emp)
    }

    for projectsRows.Next() {
        var proj string
        err := projectsRows.Scan(&proj)
        if err != nil {
            log.Println("Ошибка чтения проекта:", err)
            continue
        }
        projects = append(projects, proj)
    }

    data := map[string]interface{}{
        "Employees": employees,
        "Projects":  projects,
    }

    tmpl, err := template.ParseFiles("templates/add_task.html")
    if err != nil {
        log.Println("Ошибка парсинга шаблона:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
}

// Обработка POST-запроса для добавления задачи
func handleAddTask(w http.ResponseWriter, r *http.Request) {
    selectedEmpName := r.FormValue("employee-name")
    selectedProjName := r.FormValue("project-name")
    taskPoint := r.FormValue("task-point")
	taskStatus := r.FormValue("task-status") 
    deadline := r.FormValue("deadline")

    db, err := bd.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Сначала получаем идентификаторы по выбранным данным
    var employeeID int
    err = db.QueryRow("SELECT employee_id FROM employee WHERE username=$1", selectedEmpName).Scan(&employeeID)
    if err != nil {
        http.Error(w, "Ошибка определения идентификатора сотрудника", http.StatusBadRequest)
        return
    }

    var projectID int
    err = db.QueryRow("SELECT project_id FROM projects WHERE project_name=$1", selectedProjName).Scan(&projectID)
    if err != nil {
        http.Error(w, "Ошибка определения идентификатора проекта", http.StatusBadRequest)
        return
    }

    // Далее продолжаем выполнение операции...

    stmt, err := db.Prepare(` INSERT INTO tasks(task_point, deadline, project_id) VALUES($1, $2, $3) RETURNING task_id `)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    var newTaskID int
    err = stmt.QueryRow(taskPoint, taskStatus, deadline, projectID).Scan(&newTaskID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = db.Exec(` INSERT INTO task_employee(task_id, employee_id) VALUES ($1, $2) `, newTaskID, employeeID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/tasks/"+r.URL.Query().Get("username"), http.StatusSeeOther)
}