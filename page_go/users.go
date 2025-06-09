   package page_go

   import (
       "html/template"
       "net/http"
       "log"
       "projectSUBD/bd" // Assuming bd package is in the same module
   )

   func UsersHandler(w http.ResponseWriter, r *http.Request) {
    dbusers, err := bd.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer dbusers.Close()

    rows, err := dbusers.Query(`SELECT u.username as name, u.assigned_role as role, '/tasks/' || u.username as task_link 
    FROM user_info_view u
    left join employee e on u.id = e.id_user 
    left join task_employee te on e.employee_id = te.employee_id 
    group by u.username, u.assigned_role`)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var userInfo []bd.UserInfo
    for rows.Next() {
        var u bd.UserInfo
        err := rows.Scan(&u.Name, &u.Roles, &u.TasksLink)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        userInfo = append(userInfo, u)
    }

    var tmpl = template.Must(template.ParseFiles("templates/users.html"))
    tmpl.Execute(w, userInfo)
}
