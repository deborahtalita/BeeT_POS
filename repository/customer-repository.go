package repository

import (
	"beet_pos/structs"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	InsertCustomer(customer structs.Customer) structs.Customer
	IsDuplicate(customer_phone string)(tx *gorm.DB)
}

type customerConnection struct{
	connection *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository{
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) InsertCustomer(customer structs.Customer) structs.Customer{
	customer.Customer_update = time.Now()
	db.connection.Save(&customer)
	return customer
}

func (db *customerConnection) IsDuplicate(customer_phone string)(tx *gorm.DB){
	var customer structs.Customer
	return db.connection.Where("customer_phone = ?", customer).Take(&customer)
}



