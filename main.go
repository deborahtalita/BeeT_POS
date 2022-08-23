package main

import (
	"beet_pos/config"
	"beet_pos/controllers"
	"beet_pos/middleware"
	"beet_pos/repository"
	"beet_pos/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.SetUpDatabaseConnection()
	customerRepository 	repository.CustomerRepository 	= repository.NewCustomerRepository(db)
	userRepository		repository.UserRepository 		= repository.NewUserRepository(db)
	userService			service.UserService				= service.NewUserService(userRepository)
	OutletRepository 	repository.OutletRepository 	= repository.NewOutletRepository(db)
	customerService 	service.CustomerService 		= service.NewAuthService(customerRepository)
	outletService 		service.OutletService 			= service.NewOutletService(OutletRepository)
	jwtService			service.JWTService				= service.NewJWTService()
	customerController 	controllers.CustomerController	= controllers.NewCustomerController(customerService,jwtService)
	userController		controllers.UserController		= controllers.NewuserController(userService, jwtService)
	outletController	controllers.OutletController	= controllers.NewOutletController(outletService, jwtService)
)

func main() {
	defer config.SetUpDatabaseConnection()
	// config.GetAll(db)
	router:= gin.Default()
	customerRoutes := router.Group("api/customer")
	{
		customerRoutes.POST("/register",customerController.Register)
	}
	outletRoutes := router.Group("api/outlet", middleware.AuthorizeJWT(jwtService))
	{
		outletRoutes.GET("/", outletController.GetAllOutlets)
		outletRoutes.GET("/filter", outletController.GetPaginateFiltering)
		outletRoutes.POST("/create", outletController.CreateOutlet)
		outletRoutes.GET("/read", outletController.ReadOutlet)
		outletRoutes.PUT("/update",outletController.UpdateOutlet)
		outletRoutes.GET("/read/:outlet_id", outletController.FindByID)
		outletRoutes.PATCH("/delete/:outlet_id", outletController.DeleteOutlet)
	}
	userRouters := router.Group("api/user")
	{
		userRouters.POST("/create",userController.Register)
		userRouters.GET("/", userController.ReadUser)
	}
	router.POST("/login", userController.Login)


	router.Run()
}
