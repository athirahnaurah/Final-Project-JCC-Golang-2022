package models

type User struct {
	ID        int      `json:"id" gorm:"primary_key"`
	Name      string   `json:"name" gorm:"not null;unique"`
	Email     string   `json:"email" gorm:"not null;unique"`
	Password  string   `json:"password" gorm:"not null"`
	Address   string   `json:"address"`
	Telephone string   `json:"telephone" gorm:"not null;unique"`
	Reviews   []Review `json:"-"`
	Orders    []Order  `json:"-"`
}