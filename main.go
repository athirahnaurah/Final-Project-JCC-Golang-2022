package main

import (
	"Final-Project-JCC-Golang-2022/config"
	"Final-Project-JCC-Golang-2022/docs"
	"Final-Project-JCC-Golang-2022/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger E-Commerce API"
	docs.SwaggerInfo.Description = "This is a sample server E-Commerce."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
 	r.Run()
}