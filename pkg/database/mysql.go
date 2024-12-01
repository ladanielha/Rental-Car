package database

import (
	"Rental-car/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMysql() *gorm.DB {
	db_USER := os.Getenv("DB_USER")
	db_PASSWORD := os.Getenv("DB_PASSWORD")
	db_HOST := os.Getenv("DB_HOST")
	db_PORT := os.Getenv("DB_PORT")
	db_DBNAME := os.Getenv("DB_DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_USER, db_PASSWORD, db_HOST, db_PORT, db_DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return nil
	}
	db.AutoMigrate(&models.Users{}, &models.Cars{}, &models.Booking{})
	return db
}
