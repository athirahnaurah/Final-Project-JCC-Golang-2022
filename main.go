package main

import (
	"Final-Project-JCC-Golang-2022/config"
)

func main() {
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}