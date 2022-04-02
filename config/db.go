package config

import (
	"Final-Project-JCC-Golang-2022/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
  func ConnectDataBase() *gorm.DB {
	username := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	database := "database_ecommerce"
  
	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  
	if err != nil {
	  panic(err.Error())
	}
  
	db.AutoMigrate(&models.Order{}, &models.Method{}, &models.Payment{}, &models.Category{}, &models.Product{},&models.Review{},&models.User{},&models.Admin{})
  
	return db
  }