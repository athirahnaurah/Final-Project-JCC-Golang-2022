package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Final-Project-JCC-Golang-2022/controllers"
	"Final-Project-JCC-Golang-2022/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register-admin",controllers.Register)
    r.POST("/login-admin",controllers.Login)

    r.POST("/register-customer",controllers.RegisterCustomer)
    r.POST("/login-customer",controllers.LoginCustomer)

    r.GET("/product-categories", controllers.GetAllCategory)
    r.GET("/product-categories/:id", controllers.GetProductCategoryById)
    r.GET("/product-categories/:id/product",controllers.GetProductsByCategoryId)

    productCategoryMiddlewareRoute := r.Group("/product-categories")
    productCategoryMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    productCategoryMiddlewareRoute.POST("/", controllers.CreateProductCategory)
    productCategoryMiddlewareRoute.PATCH("/:id", controllers.UpdateProductCategory)
    productCategoryMiddlewareRoute.DELETE("/:id", controllers.DeleteProductCategory)

    r.GET("/products", controllers.GetAllProducts)
    r.GET("/product/:id", controllers.GetProductById)

    productMiddlewareRoute := r.Group("/product")
    productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    productMiddlewareRoute.POST("/", controllers.CreateProduct)
    productMiddlewareRoute.PATCH("/:id", controllers.UpdateProduct)
    productMiddlewareRoute.DELETE("/:id", controllers.DeleteProduct)

    r.GET("/payment-method", controllers.GetAllPaymentMethod)
    r.GET("/payment-method/:id", controllers.GetPaymentMethodById)
    r.GET("/payment-method/:id/payments",controllers.GetPaymentsByPaymentMethodId)

    paymentMethodMiddlewareRoute := r.Group("/payment-method")
    paymentMethodMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    paymentMethodMiddlewareRoute.POST("/", controllers.CreatePaymentMethod)
    paymentMethodMiddlewareRoute.PATCH("/:id", controllers.UpdatePaymentMethod)
    paymentMethodMiddlewareRoute.DELETE("/:id", controllers.DeletePaymentMethod)

    paymentMiddlewareRoute := r.Group("/payment")
    paymentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    paymentMiddlewareRoute.GET("/", controllers.GetAllPayments)
    paymentMiddlewareRoute.GET("/:id", controllers.GetPaymentById)
    paymentMiddlewareRoute.POST("/", controllers.CreatePayment)
    paymentMiddlewareRoute.PATCH("/:id", controllers.UpdatePayment)
    paymentMiddlewareRoute.DELETE("/:id", controllers.DeletePayment)


    orderMiddlewareRoute := r.Group("/orders")
    orderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    orderMiddlewareRoute.GET("/", controllers.GetAllOrders)
    orderMiddlewareRoute.GET("/:id", controllers.GetOrderById)
    orderMiddlewareRoute.POST("/", controllers.CreateOrder)
    orderMiddlewareRoute.PATCH("/:id", controllers.UpdateOrder)
    orderMiddlewareRoute.DELETE("/:id", controllers.DeleteOrder)

    r.GET("/reviews", controllers.GetAllReviews)
    r.GET("/reviews/:id", controllers.GetReviewById)
    reviewMiddlewareRoute := r.Group("/reviews")
    reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    reviewMiddlewareRoute.POST("/", controllers.CreateReview)
    reviewMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
    reviewMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}