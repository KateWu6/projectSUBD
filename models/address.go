package models

type Address struct {
    AddressID int    `json:"address_id" gorm:"primary_key"`
    Country   string `json:"country"`
    City      string `json:"city"`
    Street    string `json:"street"`
    House     string `json:"house"`
    Flat      string `json:"flat"`
}