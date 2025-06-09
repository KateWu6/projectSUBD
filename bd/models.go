package bd

import "time"

type User struct { 
Username string 
HashedPassword []byte }

type UserInfo struct { 
Name string `json:"username"` 
Roles string `json:"roles" `
TasksLink string 
}

// Tasks represents a task associated with a project. type Tasks struct { TasksID int json:"tasks_id" TaskPoint string json:"task_point" TaskStatus string json:"task_status" Deadline time.Time json:"deadline" ProjectName string json:"project_id" }

type TasksProject struct { 
TaskID int `json:"task_id"` 
TaskPoint string `json:"task_point"` 
TaskStatus string `json:"task_status"` 
Deadline time.Time `json:"deadline"` 
EmployeeName string `json:"employee_name"` // Имя работника 
}

type TaskForm struct { 
TaskPoint string 
TaskStatus  string
Deadline string 
EmployeeID string 
ProjectName string 
}

//для распределение прав на страницы среди пользователей с помощью ролей
type PageData struct {
    Role      string
    Projects  []struct{ Title string } // Структуры с данными проектов
    Tasks     []struct{ Title string } // Структуры с данными задач
    Username  string                   // Имя пользователя
}

// Rol represents a role in the system.
type Rol struct {
 RolID   int    `json:"rol_id"`
 RolName string `json:"rol_name"`
 RolPoint string `json:"rol_point"`
}

// Address represents a physical address.
type Address struct {
 AddressID int    `json:"addres_id"`
 Country   string `json:"country"`
 City      string `json:"city"`
 Street    string `json:"street"`
 House     string `json:"house"`
 Flat      string `json:"flat"`
}

// Tasks represents a task associated with a project.
type Tasks struct {
 TasksID     int       `json:"tasks_id"`
 TaskPoint   string    `json:"task_point"`
 TaskStatus  string    `json:"task_status"`
 Deadline    time.Time `json:"deadline"`
 ProjectName   string      `json:"project_id"`
}

// StructuralDivisionsCompany represents a structural division within a company.
type StructuralDivisionsCompany struct {
 StructuralDivisionsID   int    `json:"structural_divisions_id"`
 StructuralDivisionsName string `json:"structural_divisions_name"`
 AddressID               int    `json:"address_id"`
}

// Projects represents a project.
type Projects struct {
 ProjectID   int       `json:"project_id"`
 ProjectName string    `json:"project_name"`
 ProjectGoal string    `json:"project_goal"`
 ProjectStatus string    `json:"project_status"`
 Deadlines   time.Time `json:"deadlines"`
}

// Employee represents an employee.
type Employee struct {
 EmployeeID          int    `json:"employee_id"`
 Surname             string `json:"surname"`
 Firstname           string `json:"firstname"`
 Middlename          string `json:"middlename,omitempty"` // Use omitempty if it can be null
 PassportSeries      string `json:"passport_series"`     // Using string for char
 PassportNumber      string `json:"passport_number"`     // Using string for char
 Post                string `json:"post"`
 Contacts            string `json:"contacts"`
 Photo               string `json:"photo"`
 RolID               int    `json:"rol_id"`
 AddressID           int    `json:"address_id"`
 StructuralDivisionsID int    `json:"structural_divisions_id"`
}

// TaskEmployee represents the relationship between tasks and employees.
type TaskEmployee struct {
 TaskEmployeeID int `json:"task_employee_id"`
 EmployeeID     int `json:"employee_id"`
 TaskID         int `json:"task_id"`
}

// ProjectEmployee represents the relationship between projects and employees.
type ProjectEmployee struct {
 ProjectEmployeeID int `json:"project_employee_id"`
 EmployeeID        int `json:"employee_id"`
 ProjectID         int `json:"project_id"`
}
