package models

import "time"

type Task struct {
    TaskID      int       `json:"task_id" gorm:"primary_key"`
    TaskPoint   string    `json:"task_point"`
    TaskStatus  string    `json:"task_status"`
    Deadline    time.Time `json:"deadline"`
    ProjectID   int       `json:"project_id"`
}