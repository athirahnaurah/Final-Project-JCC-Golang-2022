package controllers

import (
	"Final-Project-JCC-Golang-2022/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type paymentInput struct {
	Total float32   `json:"total"`
	Status string   `json:"status"`
	Method_id int      `json:"method_id"`
}

// GetAllPayments godoc
// @Summary Get all Payments.
// @Description Get a list of Payments.
// @Tags Payment
// @Produce json
// @Success 200 {object} []models.Payment
// @Router /payments [get]
func GetAllPayments(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var payments []models.Payment
    db.Find(&payments)

    c.JSON(http.StatusOK, gin.H{"data": payments})
}

// CreatePayments godoc
// @Summary Create New Payments.
// @Description Creating a new Payments.
// @Tags Payment
// @Param Body body paymentInput true "the body to create a new Payment"
// @Produce json
// @Success 200 {object} models.Payment
// @Router /payment [post]
func CreatePayment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
    // Validate input
    var input paymentInput
	var method models.Method
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := db.Where("id = ?", input.Method_id).First(&method).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category_id not found!"})
        return
    }

    // Create Payment
    payment := models.Payment{Total: input.Total, Status: input.Status, MethodId: input.Method_id}
    db.Create(&payment)

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// GetPaymentById godoc
// @Summary Get Payment.
// @Description Get an Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Success 200 {object} models.Payment
// @Router /payment/{id} [get]
func GetPaymentById(c *gin.Context) { // Get model if exist
    var payment models.Payment

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// UpdatePayment godoc
// @Summary Update Payment.
// @Description Update Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Param Body body paymentInput true "the body to update Payment"
// @Success 200 {object} models.Payment
// @Router /payment/{id} [patch]
func UpdatePayment(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var payment models.Payment
	var method models.Method
    if err := db.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input paymentInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := db.Where("id = ?", input.Method_id).First(&method).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Method_id not found!"})
        return
    }

    var updatedInput models.Payment
    updatedInput.Total= input.Total
    updatedInput.Status = input.Status
    updatedInput.MethodId = input.Method_id

    db.Model(&payment).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DeletePayment godoc
// @Summary Delete one Payment.
// @Description Delete a Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Success 200 {object} map[string]boolean
// @Router /payment/{id} [delete]
func DeletePayment(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var payment models.Payment
    if err := db.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&payment)

    c.JSON(http.StatusOK, gin.H{"data": true})
}