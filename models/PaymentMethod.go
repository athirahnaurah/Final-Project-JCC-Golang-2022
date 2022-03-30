package models

type Method struct {
	ID                int       `json:"id" gorm:"primary_key"`
	PaymentMethodName string    `json:"method_name"`
	Payments          []Payment `json:"-"`
}