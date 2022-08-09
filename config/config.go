package config

import (
	"beet_pos/structs"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetUpDatabaseConnection membuat new connection ke database
func SetUpDatabaseConnection() *gorm.DB{
	err := godotenv.Load()

	if err != nil{
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser,dbHost,dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to create a connection to database")
	}
	//model
	db.AutoMigrate(&structs.Outlet{},&structs.Customer{})
	db.Preload("Outlet").Find(&structs.Customer{})
	return db
}

//CloseConnectionDatabase method untuk menutup koneksi ke database
func CloseConnectionDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil{
		panic("Failed to close connection from database")
	}

	dbSQL.Close()

}
