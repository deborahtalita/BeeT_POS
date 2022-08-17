package repository

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"fmt"
	"math"

	//"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product entity.Product) entity.Product
	Update(id string, product entity.Product) entity.Product
	Delete(id string)
	GetAll() ([]entity.Product, error)
	GetAllPaginate(pagination dto.Pagination) dto.Pagination
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

// AddProduct implements ProductRepository
func (db *productConnection) AddProduct(product entity.Product) entity.Product {
	//product.Product_update = time.Now()
	db.connection.Save(&product)
	return product
}

func (db *productConnection) Update(id string, product entity.Product) entity.Product {
	//product.Product_update = time.Now()
	var upProduct = entity.Product{}
	err := db.connection.Table("products").Where("product_id = ?", id).First(&upProduct).Updates(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepo.Update] error execute query %v \n", err)
		//return nil
	}
	// db.connection.Save(&product)
	// db.connection.Model(&product).Updates()
	return product
}

func (db *productConnection) Delete(id string) {
	err := db.connection.Table("products").Where("product_id = ?", id).Update("product_status", false).Error
	if err != nil {
		fmt.Printf("[ProductRepo.Delete] error execute query %v \n", err)
		//return nil
	}
}

func (db *productConnection) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := db.connection.Find(&products).Error
	if err != nil {
		fmt.Printf("[ProductRepo.GetAll] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data")
	}
	return products, nil
}

// GetAllPaginate implements ProductRepository
func (db *productConnection) GetAllPaginate(pagination dto.Pagination) dto.Pagination{
	var pgn dto.Pagination

	totRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	totalRows := int64(totRows)

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset, and order
	var product entity.Product
	var products []entity.Product
	find := db.connection.Limit(pagination.Limit).Offset(offset).Preload("Outlet")

	find = find.Find(&products)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		fmt.Printf("[ProductRepo.GetAll] error execute query %v \n", errFind)
		return pgn
	}

	pagination.Rows = products

	// count all data
	errCount := db.connection.Model(product).Count(&totalRows).Error

	if errCount != nil {
		fmt.Printf("[ProductRepo.GetAll] error execute query %v \n", errCount)
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