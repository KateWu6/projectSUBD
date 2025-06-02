package models

type StructuralDivision struct {
    StructuralDivisionsID int    `json:"structural_divisions_id" gorm:"primary_key"`
    StructuralDivisionsName string `json:"structural_divisions_name"`
    AddressID              int    `json:"address_id"`
}