package models

import "time"

type Project struct {
    ProjectID     int       `json:"project_id" gorm:"primary_key"`
    ProjectName   string    `json:"project_name"`
    ProjectGoal   string    `json:"project_goal"`
    ProjectStatus string    `json:"project_status"`
    Deadlines     time.Time `json:"deadlines"`
}