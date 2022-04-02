package models

type Payment struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Status   string  `json:"status"`
	MethodId int     `json:"method_id"`
	Orders   []Order `json:"-"`
	Method   Method  `json:"-"`
}