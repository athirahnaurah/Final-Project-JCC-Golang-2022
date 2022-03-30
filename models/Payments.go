package models

type Payment struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Total    float32 `json:"total"`
	Status   string  `json:"status"`
	MethodId int     `json:"method_id"`
	Orders   []Order `json:"-"`
	Method   Method  `json:"-"`
}