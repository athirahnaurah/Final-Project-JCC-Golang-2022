package controllers

import (
	"net/http"

	"Final-Project-JCC-Golang-2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentMethodInput struct {
    Name        string `json:"method_name"`
}

// GetAllPaymentMethod godoc
// @Summary Get all PaymentMethod.
// @Description Get a list of PaymentMethod.
// @Tags PaymentMethod
// @Produce json
// @Success 200 {object} []models.Method
// @Router /payment-method [get]
func GetAllPaymentMethod(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var methods []models.Method
    db.Find(&methods)

    c.JSON(http.StatusOK, gin.H{"data": methods})
}

// CreatePaymentMethod godoc
// @Summary Create New Method.
// @Description Creating a new Method.
// @Tags PaymentMethod
// @Param Body body paymentMethodInput true "the body to create a new Method"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Method
// @Router /payment-method [post]
func CreatePaymentMethod(c *gin.Context) {
    // Validate input
    var input paymentMethodInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Payment
    method := models.Method{PaymentMethodName: input.Name}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&method)

    c.JSON(http.StatusOK, gin.H{"data": method})
}

// GetPaymentMethodById godoc
// @Summary Get PaymentMethod.
// @Description Get an PaymentMethod by id.
// @Tags PaymentMethod
// @Produce json
// @Param id path string true "Method id"
// @Success 200 {object} models.Method
// @Router /payment-method/{id} [get]
func GetPaymentMethodById(c *gin.Context) { // Get model if exist
    var method models.Method

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&method).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": method})
}

// GetPaymentByPaymentMethodId godoc
// @Summary Get Payments.
// @Description Get all Payments by MethodId.
// @Tags PaymentMethod
// @Produce json
// @Param id path string true "Method id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Payment
// @Router /payment-method/{id}/payments [get]
func GetPaymentsByPaymentMethodId(c *gin.Context) { // Get model if exist
    var payment []models.Payment

    db := c.MustGet("db").(*gorm.DB)

    if err := db.Where("method_id = ?", c.Param("id")).Find(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// UpdatePaymentMethod godoc
// @Summary Update Method.
// @Description Update Method by id.
// @Tags PaymentMethod
// @Produce json
// @Param id path string true "Method id"
// @Param Body body paymentMethodInput true "the body to update payment method"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Method
// @Router /payment-method/{id} [patch]
func UpdatePaymentMethod(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var method models.Method
    if err := db.Where("id = ?", c.Param("id")).First(&method).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input paymentMethodInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Method
    updatedInput.PaymentMethodName= input.Name

    db.Model(&method).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": method})
}

// DeletePaymentMethod godoc
// @Summary Delete one Method.
// @Description Delete a Method by id.
// @Tags PaymentMethod
// @Produce json
// @Param id path string true "Method id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /payment-method/{id} [delete]
func DeletePaymentMethod(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var method models.Method
    if err := db.Where("id = ?", c.Param("id")).First(&method).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&method)

    c.JSON(http.StatusOK, gin.H{"data": true})
}	