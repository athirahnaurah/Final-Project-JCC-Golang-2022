package main

import (
	"Final-Project-JCC-Golang-2022/config"
	"Final-Project-JCC-Golang-2022/docs"
	"Final-Project-JCC-Golang-2022/routes"
	"Final-Project-JCC-Golang-2022/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email naurathirahh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	 // for load godotenv
	 environment := utils.Getenv("ENVIRONMENT", "development")

	 if environment == "development" {
	   err := godotenv.Load()
 
	   if err != nil {
		 log.Fatal("Error loading .env file")
	   }
	 }
 
	
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger E-Commerce API"
	docs.SwaggerInfo.Description = "This is a sample server E-Commerce."
	docs.SwaggerInfo.InfoInstanceName = "Athirah Naurah F"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
 	r.Run()
}