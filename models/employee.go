package models

type Employee struct {
    EmployeeID             int    `json:"employee_id" gorm:"primary_key"`
    Surname                string `json:"surname"`
    Firstname              string `json:"firstname"`
    Middlename             string `json:"middlename"`
    PassportSeries         string `json:"passport_series"`
    PassportNumber         string `json:"passport_number"`
    Post                   string `json:"post"`
    Contacts               string `json:"contacts"`
    Photo                  string `json:"photo"`
    RolID                  int    `json:"rol_id"`
    AddressID              int    `json:"address_id"`
    StructuralDivisionsID  int    `json:"structural_divisions_id"`
}