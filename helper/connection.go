package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("cannot load environment %v", err)
	}

	login := os.Getenv("login")
	db, errData := gorm.Open(mysql.Open(login), &gorm.Config{})
	if errData != nil {
		log.Fatalf("cannot connect database %v", errData)
	}

	DB = db

	return db
}
