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
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetAllProds(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductController(productService service.ProductService, jwtService service.JWTService) ProductController {
	return &productController{
		productService: productService,
		jwtService:     jwtService,
	}
}

func (c *productController) AddProduct(ctx *gin.Context) {
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
		outletID := fmt.Sprintf("%v", claims["Outlet_id"])
		log.Printf(outletID)
		addProductDTO.Outlet_id = outletID
		createdProduct := c.productService.AddProduct(addProductDTO)
		response := helper.BuildResponse(true, "OK!", createdProduct)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *productController) Update(ctx *gin.Context) {
	id := ctx.Param("product_id")
	var productUpdateDTO dto.UpdateProductDTO
	errDTO := ctx.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	res := c.productService.Update(id, productUpdateDTO)
	response := helper.BuildResponse(true, "Update successful!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) Delete(ctx *gin.Context) {
	id := ctx.Param("product_id")

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	c.productService.Delete(id)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetAll(ctx *gin.Context) {
	products, err := c.productService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	res := helper.BuildResponse(true, "OK!", products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetAllProds(ctx *gin.Context){
	pagination := helper.GeneratePagination(ctx)
	products := c.productService.GetAllPaginate(*pagination)
	res := helper.BuildResponse(true, "OK", products)
	ctx.JSON(http.StatusOK, res)
}