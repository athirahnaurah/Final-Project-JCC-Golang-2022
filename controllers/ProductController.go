package controllers

import (
	"Final-Project-JCC-Golang-2022/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	ProductName string   `json:"product"`
	Description string   `json:"desc"`
	Category_id int      `json:"category_id"`
}

// GetAllProducts godoc
// @Summary Get all Products.
// @Description Get a list of Products.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var products []models.Product
    db.Find(&products)

    c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body productInput true "the body to create a new Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
    // Validate input
    var input productInput
	var category models.Category
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := db.Where("id = ?", input.Category_id).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category_id not found!"})
        return
    }

    // Create Rating
    product := models.Product{ProductName: input.ProductName, Description: input.Description, Category_id: input.Category_id}
    db.Create(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductById godoc
// @Summary Get Product.
// @Description Get an Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
    var product models.Product

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update Product"
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var product models.Product
	var category models.Category
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := db.Where("id = ?", input.Category_id).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category_id not found!"})
        return
    }

    var updatedInput models.Product
    updatedInput.ProductName= input.ProductName
    updatedInput.Description = input.Description
    updatedInput.Category_id = input.Category_id

    db.Model(&product).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var product models.Product
    if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&product)

    c.JSON(http.StatusOK, gin.H{"data": true})
}