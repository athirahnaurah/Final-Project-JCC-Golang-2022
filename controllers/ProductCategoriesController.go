package controllers

import (
	"net/http"

	"Final-Project-JCC-Golang-2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productCategoryInput struct {
    CategoryName        string `json:"category"`
}

// GetAllProductCategories godoc
// @Summary Get all ProductCategories.
// @Description Get a list of ProductCategories.
// @Tags ProductCategory
// @Produce json
// @Success 200 {object} []models.Category
// @Router /product-categories [get]
func GetAllCategory(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var categories []models.Category
    db.Find(&categories)

    c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateProductCategory godoc
// @Summary Create New ProductCategory.
// @Description Creating a new ProductCategory.
// @Tags ProductCategory
// @Param Body body productCategoryInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /product-categories [post]
func CreateProductCategory(c *gin.Context) {
    // Validate input
    var input productCategoryInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Category
    category := models.Category{CategoryName: input.CategoryName}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&category)

    c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetProductCategoryById godoc
// @Summary Get ProductCategory.
// @Description Get an ProductCategory by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /product-categories/{id} [get]
func GetProductCategoryById(c *gin.Context) { // Get model if exist
    var category models.Category

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetProductByProductCategoryId godoc
// @Summary Get Products.
// @Description Get all Product by ProductCategoryId.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} []models.Product
// @Router /product-categories/{id}/product [get]
func GetProductsByCategoryId(c *gin.Context) { // Get model if exist
    var products []models.Product

    db := c.MustGet("db").(*gorm.DB)

    if err := db.Where("category_id = ?", c.Param("id")).Find(&products).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": products})
}

// UpdateProductCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "Category id"
// @Param Body body productCategoryInput true "the body to update product category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Category
// @Router /product-categories/{id} [patch]
func UpdateProductCategory(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var category models.Category
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input productCategoryInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Category
    updatedInput.CategoryName = input.CategoryName

    db.Model(&category).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteProductCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags ProductCategory
// @Produce json
// @Param id path string true "Category id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /product-categories/{id} [delete]
func DeleteProductCategory(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var category models.Category
    if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&category)

    c.JSON(http.StatusOK, gin.H{"data": true})
}	