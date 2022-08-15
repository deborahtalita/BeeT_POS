package controller

import (
	"beet_pos/dto"
	"beet_pos/helper"
	"beet_pos/service"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddProduct(ctx *gin.Context)
}

type productController struct {
	productService  service.ProductService
	jwtService		service.JWTService
}

func NewProductController(productService service.ProductService, jwtService service.JWTService) ProductController {
	return &productController {
		productService: productService,
		jwtService: jwtService,
	}
}

func (c *productController) AddProduct(ctx *gin.Context){
	var addProductDTO dto.AddProductDTO
	errDTO := ctx.ShouldBind(&addProductDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		token, err := c.jwtService.ValidateToken(authHeader)
		if err != nil {
			panic(err.Error())
		}
		claims := token.Claims.(jwt.MapClaims)
		outletID := fmt.Sprintf("%v", claims["Outlet_ID"])
		log.Printf(outletID)
		addProductDTO.Outlet_id = outletID
		createdProduct := c.productService.AddProduct(addProductDTO)
		response := helper.BuildResponse(true, "OK!", createdProduct)
		ctx.JSON(http.StatusCreated, response)
	}
}

// func (c *customerController) getOutletIDByToken(token string) string {
// 	log.Printf(token)
// 	jToken, err := c.jwtService.ValidateToken(token)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	claims := jToken.Claims.(jwt.MapClaims)
// 	outlet_id := fmt.Sprintf("%v", claims["Outlet_ID"])
// 	log.Printf(outlet_id)
// 	return outlet_id
// }