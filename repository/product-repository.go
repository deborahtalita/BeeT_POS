package repository

import (
	"beet_pos/entity"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product entity.Product) entity.Product
	Update(product entity.Product) entity.Product
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
	product.Product_update = time.Now()
	db.connection.Save(&product)
	return product
}

func (db *productConnection) Update(product entity.Product) entity.Product {
	product.Product_update = time.Now()
	db.connection.Save(&product)
	return product
}
