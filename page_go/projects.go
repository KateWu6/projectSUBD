package page_go

import (
    "net/http"
	"html/template"
	"projectSUBD/bd"
	"log"
    "time"
)

func FormatDate(t time.Time) string {
    return t.Format("2006-01-02")
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
    dbproject, err := bd.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer dbproject.Close()

    rows, err := dbproject.Query("SELECT * FROM projects")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var projects []bd.Projects
    for rows.Next() {
        var p bd.Projects
        err := rows.Scan(&p.ProjectID, &p.ProjectName, &p.ProjectGoal, &p.ProjectStatus, &p.Deadlines)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        projects = append(projects, p)
    }

    var tmpl = template.Must(template.New("projects.html").Funcs(template.FuncMap{
    "FormatDate": FormatDate,
    }).ParseFiles("templates/projects.html"))
    tmpl.Execute(w, projects)
}
