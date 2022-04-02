package controllers

import (
	"Final-Project-JCC-Golang-2022/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginUserInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type RegisterUserInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email    string `json:"email" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
    Address    string `json:"address" binding:"required"`
}

// LoginCustomer godoc
// @Summary Login as as customer.
// @Description Logging in to get jwt token to access customer api by roles. User can create order, payment, and review a product.
// @Tags Customer
// @Param Body body LoginUserInput true "the body to login a customer"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login-customer [post]
func LoginCustomer(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input LoginUserInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    u := models.User{}

    u.Username = input.Username
    u.Password = input.Password

    token, err := models.LoginCheckUser(u.Username, u.Password, db)

    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
        return
    }

    customer := map[string]string{
        "username": u.Username,
        "email":    u.Email,
		"telephone" : u.Telephone,
    }

    c.JSON(http.StatusOK, gin.H{"message": "login success", "customer": customer, "token": token})

}

// Register godoc
// @Summary Register a customer.
// @Description registering a customer from public access.
// @Tags Customer
// @Param Body body RegisterUserInput true "the body to register a customer"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register-customer [post]
func RegisterCustomer(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input RegisterUserInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    u := models.User{}

    u.Username = input.Username
    u.Email = input.Email
    u.Password = input.Password
    u.Telephone = input.Telephone
    u.Address = input.Address

    _, err := u.SaveUser(db)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    customer := map[string]string{
        "username": input.Username,
        "email":    input.Email,
		"telephone" : input.Telephone,
		"address" : input.Address,
    }

    c.JSON(http.StatusOK, gin.H{"message": "registration success", "customer": customer})

}