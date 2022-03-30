package models

import "time"

type Order struct {
	ID int `json:"id" gorm:"primary_key"`
	OrderDate time.Time `json:"date"`
	Quantity int `json:"qty"`
	ProductId int `json:"product_id"`
	PaymentId int `json:"payment_id"`
	UserId int`json:"user_id"`
	Product   Product `json:"-"`
	Payment   Payment `json:"-"`
	User   User `json:"-"`
}