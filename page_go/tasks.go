package page_go

import (
    "net/http"
	"html/template"
    "log"
    "strconv"
    "projectSUBD/bd"
    "github.com/gorilla/mux"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]


    db, err := bd.Connect()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    rows, err := db.Query(` SELECT 
    t.task_id, 
    t.task_point, 
    t.task_status, 
    t.deadline, 
    p.project_name AS project_name
FROM 
    tasks t
INNER JOIN 
    task_employee te ON t.task_id = te.task_id
INNER JOIN 
    employee e ON te.employee_id = e.employee_id
INNER JOIN 
    user_info_view uv ON e.id_user = uv.id
INNER JOIN 
    projects p ON t.project_id = p.project_id -- подключение таблицы проектов
WHERE 
    uv.username = $1`, username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []bd.Tasks
    for rows.Next() {
        var task bd.Tasks
        err := rows.Scan(&task.TasksID, &task.TaskPoint, &task.TaskStatus, &task.Deadline, &task.ProjectName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

     tmpl, err := template.ParseFiles("templates/tasks.html")
    if err != nil {
        log.Println("Ошибка парсинга шаблона:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, tasks)
}

func ProjectTasksHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    projectIDStr := vars["projectID"]
    projectID, err := strconv.Atoi(projectIDStr)
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }

    db, err := bd.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query(`  SELECT t.task_id, 
    t.task_point, 
    t.task_status, 
    t.deadline, 
    uv.username AS user_name 
    FROM tasks t  
INNER JOIN 
    task_employee te ON t.task_id = te.task_id
INNER JOIN 
    employee e ON te.employee_id = e.employee_id
INNER JOIN 
    user_info_view uv ON e.id_user = uv.id
INNER JOIN 
    projects p ON t.project_id = p.project_id -- подключение таблицы проектов
    WHERE t.project_id = $1 `, projectID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []bd.TasksProject
    for rows.Next() {
        var task bd.TasksProject
        err := rows.Scan(&task.TaskID, &task.TaskPoint, &task.TaskStatus, &task.Deadline,  &task.EmployeeName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

    tmpl, err := template.New("project_tasks.html").Funcs(template.FuncMap{"FormatDate": FormatDate}).ParseFiles("templates/project_tasks.html")
    if err != nil {
        log.Println("Ошибка парсинга шаблона:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, tasks)
}