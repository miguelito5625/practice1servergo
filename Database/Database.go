package Database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "administrador:W1f1n3t.@tcp(34.69.200.220:3306)/fortesting?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE CONNECT FAILED, status: ", err)
	}
	// Migrate the schema
	// db.AutoMigrate(&Models.Product{})

	DB = db

}
