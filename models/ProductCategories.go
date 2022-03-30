package models

type Category struct {
	ID           int       `json:"id" gorm:"primary_key"`
	CategoryName string    `json:"category"`
	Products     []Product `json:"-"`
}