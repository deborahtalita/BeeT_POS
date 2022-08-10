package config

import (
	"beet_pos/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {

	dsn := "host=localhost user=postgres password=admin dbname=BeeT_POS port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Customer{}, &entity.Outlet{})
	//db.Preload("Outlet").Find(&entity.Customer{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}