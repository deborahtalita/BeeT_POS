package main

import (
	"beet_pos/config"
	"beet_pos/controller"
	"beet_pos/middleware"
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
	productRepository	repository.ProductRepository	= repository.NewProductRepository(db)
	userService			service.UserService				= service.NewUserService(userRepository)
	customerService 	service.CustomerService			= service.NewAuthService(customerRepository)
	productService		service.ProductService			= service.NewProductService(productRepository)
	jwtService			service.JWTService				= service.NewJWTService()
	userController		controller.UserController		= controller.NewuserController(userService, jwtService)
	customerController	controller.CustomerController	= controller.NewCustomerController(customerService, jwtService)
	productController	controller.ProductController	= controller.NewProductController(productService, jwtService)
)

func main() {
	router := gin.Default()

	customerRoutes := router.Group("api/customer",middleware.AuthorizeJWT(jwtService))
	{
		customerRoutes.POST("/register",customerController.Register)
	}
	productRoutes := router.Group("api/products",middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.POST("/", productController.AddProduct)
		productRoutes.GET("/", productController.GetAllProds)
		productRoutes.GET("/:product_id", productController.GetProductByID)
		productRoutes.PATCH("/:product_id", productController.Update)
		productRoutes.PATCH("/delete/:product_id",productController.Delete)
	}
	productVariantRoutes := router.Group("api/productvariants",middleware.AuthorizeJWT(jwtService))
	{
		productVariantRoutes.POST("/:product_id", productController.AddVariant)
	}

	productDiscountRoutes := router.Group("api/productdiscounts",middleware.AuthorizeJWT(jwtService))
	{
		productDiscountRoutes.POST("/:product_id", productController.AddDiscount)
	}
	productPictureRoutes := router.Group("api/productpictures",middleware.AuthorizeJWT(jwtService))
	{
		productPictureRoutes.POST("/:product_id",productController.AddPicture)
	}
	router.POST("/login", userController.Login)
	router.POST("/auth/refresh",middleware.AuthorizeJWT(jwtService),userController.Refresh)
	log.Fatal(router.Run(":8080"))
	
	router.Run()
}