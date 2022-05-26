package database

import (
	"log"

	"go-jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/gojwt"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("tidak dapat terkoneksi dengan database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
