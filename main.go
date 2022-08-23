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
	db *gorm.DB	= config.SetupDatabaseConnection()
	userRepository		repository.UserRepository 	= repository.NewUserRepository(db)
	customerRepository	repository.CustomerRepository	= repository.NewCustomerRepository(db)
	productRepository	repository.ProductRepository	= repository.NewProductRepository(db)
	OutletRepository 	repository.OutletRepository 	= repository.NewOutletRepository(db)

	userService		service.UserService		= service.NewUserService(userRepository)
	customerService 	service.CustomerService		= service.NewAuthService(customerRepository)
	productService		service.ProductService		= service.NewProductService(productRepository)
	outletService 		service.OutletService 		= service.NewOutletService(OutletRepository)
	jwtService		service.JWTService		= service.NewJWTService()
	userController		controller.UserController	= controller.NewuserController(userService, jwtService)
	customerController	controller.CustomerController	= controller.NewCustomerController(customerService, jwtService)
	productController	controller.ProductController	= controller.NewProductController(productService, jwtService)
	outletController	controller.OutletController	= controller.NewOutletController(outletService, jwtService)

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
	//router.POST("/auth/refresh",middleware.AuthorizeJWT(jwtService),userController.Refresh)
	log.Fatal(router.Run(":8080"))
	
	router.Run()
}
