package models

type Role struct {
    RolID   int    `json:"rol_id" gorm:"primary_key"`
    RolName string `json:"rol_name"`
    RolPoint string `json:"rol_point"`
}