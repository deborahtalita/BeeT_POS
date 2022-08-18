package repository

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type OutletRepository interface {
	InsertOutlet(outlet entity.Outlet) entity.Outlet
	DeleteOutlet(outlet entity.Outlet) error
	ReadOutlet() [] entity.Outlet
	FindByID(outlet_id string) entity.Outlet
	UpdateOutlet(outlet entity.Outlet) entity.Outlet
	IsDuplicate(outlet_name string) (tx *gorm.DB)
	GetAllPaginate(pagination dto.Pagination) dto.Pagination
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
	outlet.Outlet_created = formatTime()
	outlet.Outlet_update = formatTime()
	db.connection.Save(&outlet)
	return outlet
}

func (db *outletConnection) UpdateOutlet(outlet entity.Outlet) entity.Outlet{
	// timein := time.Now().Add(time.Hour *60 + time.Minute * 60 + time.Second * 3600)
	// outlet.Outlet_created = oldTime(outlet.Outlet_created)
	outlet.Outlet_update = formatTime()
	db.connection.Save(&outlet)
	return outlet
}

func (db *outletConnection) IsDuplicate(outlet_name string)(tx *gorm.DB){
	var outlet entity.Outlet
	return db.connection.Where("outlet_name = ?", outlet_name).Take(&outlet)
}

func (db *outletConnection) ReadOutlet() [] entity.Outlet{
	// var outlet entity.Outlet
	// db.connection.Find(&outlet, outletID)
	// return outlet
	var outlet []entity.Outlet
	db.connection.Find(&outlet)
	return outlet
}

func (db *outletConnection) FindByID(outlet_id string) entity.Outlet{
	var outlet entity.Outlet
	db.connection.Where("outlet_id = ? ", outlet_id).Find(&outlet)
	return outlet
}

func (db *outletConnection) DeleteOutlet(outlet entity.Outlet) error{
	if err := db.connection.Delete(&outlet); err!= nil  {
		return err.Error
	}
	return nil
}

func formatTime() string{
	currentTime := time.Now()

	currentYear := currentTime.Year()
	currentMonth := currentTime.Month()
	currentDay := currentTime.Day()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()
	currentSecond := currentTime.Second()

	formatTime := fmt.Sprintf("%v-%v-%v %v:%v:%v", currentDay, currentMonth, currentYear, currentHour, currentMinute, currentSecond) 

	return formatTime

}


// func oldTime(olddate string) string{
// 	var outlet entity.Outlet
// 	olddate = outlet.Outlet_created
// 	return olddate
// }

// GetAllPaginate implements OutletsRepository
func (db *outletConnection) GetAllPaginate(pagination dto.Pagination) dto.Pagination{
	var pgn dto.Pagination

	totRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	totalRows := int64(totRows)

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset, and order
	var outlet entity.Outlet
	var outlets []entity.Outlet
	find := db.connection.Limit(pagination.Limit).Offset(offset)
	// .Preload("Outlet")

	find = find.Find(&outlets)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		fmt.Printf("[OutletRepo.GetAll] error execute query %v \n", errFind)
		return pgn
	}

	pagination.Rows = outlets

	// count all data
	errCount := db.connection.Model(outlet).Count(&totalRows).Error

	if errCount != nil {
		fmt.Printf("[OutletRepo.GetAll] error execute query %v \n", errCount)
		return pagination
	}

	pagination.TotalRows = totalRows

	//calculate total pages

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1
	fmt.Printf("totalpages: %d \n",totalPages)

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	fmt.Printf("fromrow: %d \n",fromRow)
	pagination.ToRow = toRow

	return pagination
}
