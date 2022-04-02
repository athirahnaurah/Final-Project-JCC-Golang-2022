package models

import (
	"html"
	"strings"

	"Final-Project-JCC-Golang-2022/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type User struct {
	ID        int      `json:"id" gorm:"primary_key"`
	Username  string   `json:"username" gorm:"not null;unique"`
	Email     string   `json:"email" gorm:"not null;unique"`
	Password  string   `json:"password" gorm:"not null"`
	Address   string   `json:"address"`
	Telephone string   `json:"telephone" gorm:"not null;unique"`
	Reviews   []Review `json:"-"`
	Orders    []Order  `json:"-"`
}

func VerifyPasswordUser(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheckUser(username string, password string, db *gorm.DB) (string, error) {

    var err error

    u := User{}

    err = db.Model(User{}).Where("username = ?", username).Take(&u).Error

    if err != nil {
        return "", err
    }

    err = VerifyPasswordUser(password, u.Password)

    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return "", err
    }

    token, err := token.GenerateToken(uint(u.ID))

    if err != nil {
        return "", err
    }

    return token, nil

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
    //turn password into hash
    hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if errPassword != nil {
        return &User{}, errPassword
    }
    u.Password = string(hashedPassword)
    //remove spaces in username
    u.Username = html.EscapeString(strings.TrimSpace(u.Username))

    var err error = db.Create(&u).Error
    if err != nil {
        return &User{}, err
    }
    return u, nil
}