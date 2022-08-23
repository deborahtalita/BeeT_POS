package repository

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"fmt"
	"math"
	"strings"
	"time"

	"gorm.io/gorm"
)

type OutletRepository interface {
	InsertOutlet(outlet entity.Outlet) entity.Outlet
	DeleteOutlet(outlet_id string)
	ReadOutlet() [] entity.Outlet
	FindByID(outlet_id string) entity.Outlet
	UpdateOutlet(outlet entity.Outlet) entity.Outlet
	IsDuplicate(outlet_name string) (tx *gorm.DB)
	GetAllPaginate(pagination dto.Pagination) dto.Pagination
	GetPaginateFiltering(pagination dto.Pagination) dto.Pagination
	// 
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
	var outlet []entity.Outlet
	db.connection.Find(&outlet)
	return outlet
}

func (db *outletConnection) FindByID(outlet_id string) entity.Outlet{
	var outlet entity.Outlet
	db.connection.Where("outlet_id = ? ", outlet_id).Find(&outlet)
	return outlet
}

func (db *outletConnection) DeleteOutlet(outlet_id string){
	err := db.connection.Table("outlets").Where("outlet_id = ?", outlet_id).Update("outlet_status", false).Error

	if err!= nil  {
		fmt.Printf("[ProductRepo.Update] error execute query %v \n", err)

	}

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


func (db *outletConnection) GetPaginateFiltering(pagination dto.Pagination) dto.Pagination{
	var pgn dto.Pagination

	totRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	totalRows := int64(totRows)

	offset := pagination.Page * pagination.Limit
	
	// get data with limit, offset, and order
	var outlet entity.Outlet
	var outlets []entity.Outlet

	find := db.connection.Limit(pagination.Limit).Offset(offset)
	// .Preload("Outlet")

		searchs := pagination.Searchs



	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	find = find.Preload("User").Find(&outlets)

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

//===================================== ZONED =========================================================
// func (db *outletConnection) GetPaginateFiltering(pagination dto.Pagination) (RepositoryResult, int){
// 	// var pgn dto.Pagination

// 	totRows, totalPages, fromRow, toRow := 0, 0, 0, 0
// 	totalRows := int64(totRows)

// 	offset := pagination.Page * pagination.Limit
	
// 	// get data with limit, offset, and order
// 	var outlet entity.Outlet
// 	var outlets []entity.Outlet
	

// 	find := db.connection.Limit(pagination.Limit).Offset(offset).Preload("User")

// 	// .Preload("Outlet")
	
// 	searchs := pagination.Searchs



// 	if searchs != nil {
// 		for _, value := range searchs {
// 			column := value.Column
// 			action := value.Action
// 			query := value.Query

// 			switch action {
// 			case "equals":
// 				whereQuery := fmt.Sprintf("%s = ?", column)
// 				find = find.Where(whereQuery, query)
// 				break
// 			case "contains":
// 				whereQuery := fmt.Sprintf("%s LIKE ?", column)
// 				find = find.Where(whereQuery, "%"+query+"%")
// 				break
// 			case "in":
// 				whereQuery := fmt.Sprintf("%s IN (?)", column)
// 				queryArray := strings.Split(query, ",")
// 				find = find.Where(whereQuery, queryArray)
// 				break

// 			}
// 		}
// 	}


// 	find = find.Find(&outlets)

// 	// has error find data
// 	errFind := find.Error

// 	if errFind != nil {
// 		fmt.Printf("[OutletRepo.GetAll] error execute query %v \n", errFind)
// 		return RepositoryResult{Error: errFind}, totalPages
// 	}

// 	pagination.Rows = outlets

// 	// count all data
// 	errCount := db.connection.Model(outlet).Count(&totalRows).Error

// 	if errCount != nil {
// 		fmt.Printf("[OutletRepo.GetAll] error execute query %v \n", errCount)
// 		return RepositoryResult{Error: errFind}, totalPages
// 	}

// 	pagination.TotalRows = totalRows

// 	//calculate total pages

// 	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1
// 	// fmt.Printf("totalpages: %d \n",totalPages)

// 	if pagination.Page == 0 {
// 		// set from & to row on first page
// 		fromRow = 1
// 		toRow = pagination.Limit
// 	} else {
// 		if pagination.Page <= totalPages {
// 			// calculate from & to row
// 			fromRow = pagination.Page*pagination.Limit + 1
// 			toRow = (pagination.Page + 1) * pagination.Limit
// 		}
// 	}

// 	if toRow > int(totalRows) {
// 		// set to row with total rows
// 		toRow = int(totalRows)
// 	}

// 	pagination.FromRow = fromRow
// 	fmt.Printf("from row: %d \n",fromRow)
// 	pagination.ToRow = toRow

// 	return RepositoryResult{Result: pagination}, totalPages
// }
