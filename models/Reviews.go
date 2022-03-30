package models

type Review struct {
	ID        int     `json:"id" gorm:"primary_key"`
	Comment   string  `json:"comment"`
	ProductId int     `json:"product_id"`
	UserId    int     `json:"user_id"`
	Product   Product `json:"-"`
	User      User    `json:"-"`
}