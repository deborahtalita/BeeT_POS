package main

import (
	"beet_pos/config"
	"beet_pos/controller"
	"beet_pos/repository"
	"beet_pos/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db 					*gorm.DB						= config.SetupDatabaseConnection()
	userRepository		repository.UserRepository 		= repository.NewUserRepository(db)
	customerRepository	repository.CustomerRepository	= repository.NewCustomerRepository(db)
	userService			service.UserService				= service.NewUserService(userRepository)
	customerService 	service.CustomerService			= service.NewAuthService(customerRepository)
	jwtService			service.JWTService				= service.NewJWTService()
	userController		controller.UserController		= controller.NewuserController(userService, jwtService)
	customerController	controller.CustomerController	= controller.NewCustomerController(customerService, jwtService)
)

func main() {
	router := gin.Default()

	customerRoutes := router.Group("api/customer")
	{
		customerRoutes.POST("/register",customerController.Register)
	}
	router.POST("/login", userController.Login)
	log.Fatal(router.Run(":8080"))
	
	router.Run()
}