package models

type ProjectEmployee struct {
    ProjectEmployeeID int `json:"project_employee_id" gorm:"primary_key"`
    EmployeeID        int `json:"employee_id"`
    ProjectID         int `json:"project_id"`
}