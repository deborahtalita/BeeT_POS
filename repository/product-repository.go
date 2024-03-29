package repository

import (
	"beet_pos/dto"
	"beet_pos/entity"
	"fmt"
	"math"
	"strings"

	//"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product entity.Product) entity.Product
	Update(id string, product entity.Product) entity.Product
	Delete(id string)
	FindByID(id string) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	GetAllPaginate(outlet_id string, pagination dto.Pagination) dto.Pagination
	AddVariant(variant entity.Product_variant, id string) entity.Product_variant
	AddDiscount(discount entity.Product_discount, id string) entity.Product_discount
	AddPicture(picture entity.Product_picture, id string) entity.Product_picture
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
	db.connection.Preload("Outlet").Find(&product)
	return product
}

func (db *productConnection) Update(id string, product entity.Product) entity.Product {
	//product.Product_update = time.Now()
	var upProduct = entity.Product{}
	err := db.connection.Table("products").Where("product_id = ?", id).First(&upProduct).Updates(&product).Error
	// tambah preload
	if err != nil {
		fmt.Printf("[ProductRepo.Update] error execute query %v \n", err)
		//return nil
	}
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
func (db *productConnection) GetAllPaginate(outlet_id string, pagination dto.Pagination) dto.Pagination {
	var pgn dto.Pagination

	totRows, totalPages, fromRow, toRow := 0, 0, 0, 0
	totalRows := int64(totRows)

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset, and order
	var product entity.Product
	var products []entity.Product
	find := db.connection.Limit(pagination.Limit).Offset(offset).Where("outlet_id = ?", outlet_id)

	// generate where query
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

	find = find.Preload("Product_discounts").Preload("Product_variants").Find(&products)

	// has error find data
	errFind := find.Error

	if errFind != nil {
		fmt.Printf("[ProductRepo.GetAll] error execute query %v \n", errFind)
		return pgn
	}

	pagination.Rows = products

	// count all data
	errCount := db.connection.Model(product).Where("outlet_id = ?", outlet_id).Count(&totalRows).Error

	if errCount != nil {
		fmt.Printf("[ProductRepo.GetAll] error execute query %v \n", errCount)
		return pagination
	}

	pagination.TotalRows = totalRows

	//calculate total pages

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1
	fmt.Printf("totalpages: %d \n", totalPages)

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
	fmt.Printf("totalrow: %d \n", totalRows)
	pagination.ToRow = toRow

	return pagination
}

// FindByID implements ProductRepository
func (db *productConnection) FindByID(id string) (entity.Product, error) {
	var product entity.Product
	err := db.connection.Where("product_id = ?", id).Preload("Product_discounts").Preload("Product_variants").First(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepo.FindByID] error execute query %v \n", err)
		return product, err
	}
	return product, nil
}

func (db *productConnection) AddVariant(variant entity.Product_variant, id string) entity.Product_variant {
	db.connection.Save(&variant)
	return variant
}

// AddDiscount implements ProductRepository
func (db *productConnection) AddDiscount(discount entity.Product_discount, id string) entity.Product_discount {
	db.connection.Save(&discount)
	return discount
}

func (db *productConnection) AddPicture(picture entity.Product_picture, id string) entity.Product_picture {
	db.connection.Save(&picture)
	return picture
}