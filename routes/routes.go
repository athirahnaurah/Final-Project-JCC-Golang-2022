package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Final-Project-JCC-Golang-2022/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })
    r.GET("/product-categories", controllers.GetAllCategory)
    r.POST("/product-categories", controllers.CreateProductCategory)
    r.GET("/product-categories/:id", controllers.GetProductCategoryById)
    r.PATCH("/product-categories/:id", controllers.UpdateProductCategory)
    r.DELETE("product-categories/:id", controllers.DeleteProductCategory)

    // r.GET("/age-rating-categories", controllers.GetAllRating)
    // r.POST("/age-rating-categories", controllers.CreateRating)
    // r.GET("/age-rating-categories/:id", controllers.GetRatingById)
    // r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingId)
    // r.PATCH("/age-rating-categories/:id", controllers.UpdateRating)
    // r.DELETE("age-rating-categories/:id", controllers.DeleteRating)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}