package models

import (
	"html"
	"strings"

	"Final-Project-JCC-Golang-2022/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
    // User
    Admin struct {
        ID        uint      `json:"id" gorm:"primary_key"`
        Username  string    `gorm:"not null;unique" json:"username"`
        Email     string    `gorm:"not null;unique" json:"email"`
        Password  string    `gorm:"not null;" json:"password"`
    }
)

func VerifyPassword(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

    var err error

    a := Admin{}

    err = db.Model(Admin{}).Where("username = ?", username).Take(&a).Error

    if err != nil {
        return "", err
    }

    err = VerifyPassword(password, a.Password)

    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return "", err
    }

    token, err := token.GenerateToken(a.ID)

    if err != nil {
        return "", err
    }

    return token, nil

}

func (a *Admin) SaveAdmin(db *gorm.DB) (*Admin, error) {
    //turn password into hash
    hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
    if errPassword != nil {
        return &Admin{}, errPassword
    }
    a.Password = string(hashedPassword)
    //remove spaces in username
    a.Username = html.EscapeString(strings.TrimSpace(a.Username))

    var err error = db.Create(&a).Error
    if err != nil {
        return &Admin{}, err
    }
    return a, nil
}
