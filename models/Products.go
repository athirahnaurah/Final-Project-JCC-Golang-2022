package models

type Product struct {
	ID          int      `gorm:"primary_key" json:"id"`
	ProductName string   `json:"product"`
	Description string   `json:"desc"`
	Price       float32  `json:"price"`
	Category_id int      `json:"category_id"`
	Reviews     []Review `json:"-"`
	Orders      []Order  `json:"-"`
	Category    Category `json:"-"`
}
