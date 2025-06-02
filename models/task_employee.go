package models

type TaskEmployee struct {
    TaskEmployeeID int `json:"task_employee_id" gorm:"primary_key"`
    EmployeeID     int `json:"employee_id"`
    TaskID         int `json:"task_id"`
}