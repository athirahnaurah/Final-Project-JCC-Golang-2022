package controllers

import (
	"net/http"

	"Final-Project-JCC-Golang-2022/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
    Comment   string  `json:"comment"`
	ProductId int     `json:"product_id"`
	UserId    int     `json:"user_id"`
}

// GetAllReviews godoc
// @Summary Get all reviews.
// @Description Get a list of Reviews.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReviews(c *gin.Context) {
    // get db from gin context
    db := c.MustGet("db").(*gorm.DB)
    var reviews []models.Review
    db.Find(&reviews)

    c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body reviewInput true "the body to create a new review"
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    // Validate input
    var input reviewInput
    var product models.Product
	var user models.User
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

    // Create Review
    review := models.Review{Comment: input.Comment, ProductId: input.ProductId, UserId: input.UserId}
    db.Create(&review)

    c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
    var review models.Review

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Param Body body reviewInput true "the body to update an review"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [patch]
func UpdateReview(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var review models.Review
    var product models.Product
	var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input reviewInput
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

    var updatedInput models.Review
    updatedInput.Comment = input.Comment
    updatedInput.ProductId = input.ProductId
    updatedInput.UserId= input.UserId

    db.Model(&review).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete one review.
// @Description Delete a review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} map[string]boolean
// @Router /review/{id} [delete]
func DeleteReview(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var review models.Review
    if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&review)

    c.JSON(http.StatusOK, gin.H{"data": true})
}