package main

import (
	"beet_pos/config"
	"beet_pos/controllers"
	"beet_pos/repository"
	"beet_pos/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.SetUpDatabaseConnection()
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
	customerService service.CustomerService = service.NewAuthService(customerRepository)
	customerController controllers.CustomerController = controllers.NewCustomerController(customerService)
)

func main() {
	defer config.SetUpDatabaseConnection()
	r := gin.Default()
	customerRoutes := r.Group("api/customer")
	{
		customerRoutes.POST("/register",customerController.Register)
	}

	r.Run()
}
