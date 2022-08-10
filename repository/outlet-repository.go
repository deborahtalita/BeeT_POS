package repository

import (
	"beet_pos/entity"
	"time"

	"gorm.io/gorm"
)

type OutletRepository interface {
	InsertOutlet(outlet entity.Outlet) entity.Outlet
	IsDuplicate(outlet_name string) (tx *gorm.DB)
}

type outletConnection struct{
	connection *gorm.DB
}

func NewOutletRepository(db *gorm.DB) OutletRepository{
	return &outletConnection{
		connection: db,
	}
}

func (db *outletConnection) InsertOutlet(outlet entity.Outlet) entity.Outlet{
	outlet.Outlet_created = time.Now()
	outlet.Outlet_update = time.Now()
	db.connection.Save(&outlet)
	return outlet
}

func (db *outletConnection) IsDuplicate(outlet_name string)(tx *gorm.DB){
	var outlet entity.Outlet
	return db.connection.Where("outlet_name = ?", outlet_name).Take(&outlet)
}
