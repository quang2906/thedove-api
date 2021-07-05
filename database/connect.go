package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ocg.com/model"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:abc123@tcp(127.0.0.1:3306)/thedove?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	} else {
		fmt.Println("Connect Successfully")
	}

	DB = database

	database.AutoMigrate(&model.Product{}, &model.User{}, &model.Cart{}, &model.Category{}, &model.Price{}, &model.Image{}, &model.Review{})
}
