package controllers

import (
	"net/http"
	"time"

	"Final-Project-JCC-Golang-2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderInput struct {
    Quantity int `json:"qty"`
	ProductId int `json:"product_id"`
	PaymentId int `json:"payment_id"`
	UserId int`json:"user_id"`
}

// GetAllOrders godoc
// @Summary Get all orders.
// @Description Get a list of Orders.
// @Tags Order
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Order
// @Router /orders [get]
func GetAllOrders(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var orders []models.Order
    db.Find(&orders)

    c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder godoc
// @Summary Create New Order.
// @Description Creating a new Order.
// @Tags Order
// @Param Body body orderInput true "the body to create a new order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Order
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    // Validate input
    var input orderInput
    var product models.Product
	var user models.User
	var payment models.Payment
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Where("id = ?", input.ProductId).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "productID not found!"})
        return
    }

	if err := db.Where("id = ?", input.UserId).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "userID not found!"})
        return
    }

	if err := db.Where("id = ?", input.PaymentId).First(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "userID not found!"})
        return
    }

    // Create Review
    order := models.Order{Quantity: input.Quantity, ProductId: input.ProductId, PaymentId: input.PaymentId, OrderDate: time.Now(), UserId: input.UserId}
    db.Create(&order)

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetOrderById godoc
// @Summary Get Order.
// @Description Get a Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func GetOrderById(c *gin.Context) { // Get model if exist
    var order models.Order

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// UpdateOrder godoc
// @Summary Update Order.
// @Description Update order by id.
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Param Body body orderInput true "the body to update an order"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Order
// @Router /orders/{id} [patch]
func UpdateOrder(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var order models.Order
    var product models.Product
	var user models.User
	var payment models.Payment
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input orderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Where("id = ?", input.ProductId).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ProductId not found!"})
        return
    }

	if err := db.Where("id = ?", input.UserId).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserId not found!"})
        return
    }

	if err := db.Where("id = ?", input.PaymentId).First(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "UserId not found!"})
        return
    }

    var updatedInput models.Order
    updatedInput.Quantity = input.Quantity
    updatedInput.ProductId = input.ProductId
    updatedInput.UserId= input.UserId
	updatedInput.PaymentId= input.PaymentId

    db.Model(&payment).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DeleteOrder godoc
// @Summary Delete one order.
// @Description Delete a order by id.
// @Tags Order
// @Produce json
// @Param id path string true "order id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var order models.Order
    if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&order)

    c.JSON(http.StatusOK, gin.H{"data": true})
}