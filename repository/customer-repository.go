package repository

import (
	"beet_pos/entity"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	InsertCustomer(customer entity.Customer) entity.Customer
	IsDuplicate(customer_phone string) (tx *gorm.DB)
}

type customerConnection struct {
	connection *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) InsertCustomer(customer entity.Customer) entity.Customer {
	customer.Customer_update = time.Now()
	db.connection.Save(&customer)
	db.connection.Preload("Outlet").Find(&customer)
	return customer
}

func (db *customerConnection) IsDuplicate(customer_phone string) (tx *gorm.DB) {
	var customer entity.Customer
	return db.connection.Where("customer_phone = ?", customer_phone).Take(&customer)
}
